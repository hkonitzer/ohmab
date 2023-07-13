package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/hkonitzer/ohmab/ent"
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/business"
	_ "github.com/hkonitzer/ohmab/ent/runtime"
	schemas "github.com/hkonitzer/ohmab/ent/schema"
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/hkonitzer/ohmab/internal/pkg/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Get logger
	logger := log.GetLoggerInstance()
	// Get the configuration
	config.Init()
	// Parse cmd line
	logger.Debug().Msgf("Starting Businesses Import")
	filename := flag.String("filename", "import.csv", "Filename of the CSV file to import")
	flag.Parse()
	file, err := os.OpenFile(*filename, os.O_RDONLY, 0644)
	if err != nil {
		logger.Fatal().Msgf("Error opening file: %v", err)
	}
	stats, err := file.Stat()
	if err != nil {
		logger.Fatal().Msgf("Error getting stats: %v", err)
	}
	// Create client
	ctx := context.TODO()
	// Authorize me
	uv := privacy.UserViewer{Role: privacy.Admin}
	uv.SetUserID("import")
	ctx = privacy.NewContext(ctx, &uv)
	client, clientError := db.CreateClient(ctx)
	if clientError != nil {
		logger.Fatal().Msgf("Error creating client: %v", clientError)
	}
	defer client.Close()

	logger.Info().Msgf("Start importing file '%s' with size %d bytes", stats.Name(), stats.Size())
	reader := csv.NewReader(file)
	reader.Comma = ';' // @TODO make configurable
	reader.TrimLeadingSpace = true

	successfulImports := 0
	skippedImports := 0

	header := true
	startTime := time.Now()
	var headersFromSchemas []string = nil
	var rowCounter = 0

	addressSchema := schemas.Address{}
	aFields := addressSchema.Fields()

	for {
		record, readErr := reader.Read()
		if record == nil {
			break
		}
		if readErr != nil {
			logger.Error().Msgf("Error while reading csv", readErr)
			break
		}
		rowCounter++
		if header {
			header = false
			headersFromSchemas_, err := createBusinessFieldsandCheckAgainstTableHeader(record)
			if err != nil {
				logger.Fatal().Msgf("Error checking table header: %v", err)
			} else {
				headersFromSchemas = headersFromSchemas_
				logger.Debug().Msgf("CSV headers successful validated")
			}
			continue
		}
		// Look into the database first
		busQuery := client.Business.Query()
		for i, field := range record {
			if field == "" {
				continue
			}
			schemaHeader := headersFromSchemas[i]
			schema := strings.Split(schemaHeader, "$")[0]
			schemaField := strings.Split(schemaHeader, "$")[1]
			switch schema {
			case "BUSINESS":
				switch strings.ToLower(schemaField) {
				case "name1":
					busQuery.Where(business.Name1(field))
				}
			case "ADDRESS":
			default:
				logger.Fatal().Msgf("Unknown schema '%s'", schema)
			}

		}
		// check if this business already exists
		b1, err := busQuery.Unique(true).WithAddresses().First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			logger.Error().Msgf("Import skipped, Query error: ", err)
			skippedImports++

			continue
		}
		if b1 != nil { // update business with another address
			logger.Debug().Msgf("Update Business: %v", b1.Name1)
			addressCreate := client.Address.Create()
			for i, field := range record {
				schemaHeader := headersFromSchemas[i]
				schema := strings.Split(schemaHeader, "$")[0]
				schemaField := strings.Split(schemaHeader, "$")[1]
				switch schema {
				case "BUSINESS":
				case "ADDRESS":
					if schemaField == "primary" { // @TODO: get field types, see line 199+ below
						b, _ := strconv.ParseBool(field)
						addressCreate.Mutation().SetPrimary(b)
					} else {
						err := addressCreate.Mutation().SetField(schemaField, field)
						if err != nil {
							logger.Fatal().Msgf("Error setting address field '%s' to '%s': %v", schemaField, field, err)
						}
					}

				default:
					logger.Fatal().Msgf("Unknown schema '%s'", schema)
				}
			}

			if b1.Edges.Addresses != nil {
				ae := false
				for _, addr := range b1.Edges.Addresses {
					s, _ := addressCreate.Mutation().Street()
					z, _ := addressCreate.Mutation().Zip()
					c, _ := addressCreate.Mutation().City()
					ae_ := (addr.Street == s) && (addr.Zip == z) && (addr.City == c)
					ae = ae_
					if ae_ {
						logger.Debug().Msgf("Address '%s, %s %s' already exists in business '%s'", s, z, c, b1.Name1)
						break
					}
				}
				if ae {
					skippedImports++
					continue
				}
			}
			addr, addrerr := addressCreate.Save(ctx)
			if addrerr != nil {
				t, _ := addressCreate.Mutation().Telephone()
				str, _ := addressCreate.Mutation().Street()
				cty, _ := addressCreate.Mutation().City()
				logger.Error().Msgf("Error saving address 'City=%v;Street=%v;Tel=%v' from row=%d: %v", cty, str, t, rowCounter, addrerr)
				skippedImports++
				continue
			}
			_, err := b1.Update().AddAddresses(addr).Save(ctx)
			if err != nil {
				logger.Error().Msgf("Error updating business '%s' from row=%d: %v", b1.Name1, rowCounter, err)
				skippedImports++
				continue
			}
			successfulImports++
		} else {
			logger.Debug().Msgf("Create Business: %v", record[0])
			businessCreate := client.Business.Create()
			addressCreate := client.Address.Create()
			for i, field := range record {
				if field == "" {
					continue
				}
				schemaHeader := headersFromSchemas[i]
				schema := strings.Split(schemaHeader, "$")[0]
				schemaField := strings.Split(schemaHeader, "$")[1]
				switch schema {
				case "BUSINESS":
					// no need for checking fieldType, only alias here for now
					err := businessCreate.Mutation().SetField(schemaField, field)
					if err != nil {
						logger.Fatal().Msgf("Error setting business field '%s' to '%s': %v", schemaField, field, err)
					}
				case "ADDRESS":
					var fieldType string
					for _, bField := range aFields {
						fieldDesc := bField.Descriptor()
						if fieldDesc.Name == schemaField {
							fieldType_ := fieldDesc.Info.Type.String()
							fieldType = fieldType_
							break
						}
					}
					switch fieldType {
					case "bool":
						b, err := strconv.ParseBool(field)
						if err != nil {
							logger.Fatal().Msgf("Error converting '%s' to bool from row=%d: %v", field, rowCounter, err)
						}
						err = addressCreate.Mutation().SetField(schemaField, b)
						if err != nil {
							logger.Fatal().Msgf("Error setting address field '%s' to '%s' from row=%d: %v", schemaField, field, rowCounter, err)
						}
					default: // string
						err := addressCreate.Mutation().SetField(schemaField, field)
						if err != nil {
							logger.Fatal().Msgf("Error setting address field '%s' to '%s' from row=%d: %v", schemaField, field, rowCounter, err)
						}
					}
				default:
					logger.Fatal().Msgf("Unknown schema '%s'", schema)
				}
			}
			addr, addrerr := addressCreate.Save(ctx)
			if addrerr != nil {
				t, _ := addressCreate.Mutation().Telephone()
				str, _ := addressCreate.Mutation().Street()
				cty, _ := addressCreate.Mutation().City()
				zip, _ := addressCreate.Mutation().Zip()
				ad := client.Address.Query().Where(address.CityEQ(cty), address.StreetEQ(str), address.TelephoneEQ(t), address.ZipEQ(zip)).OnlyX(ctx)
				logger.Info().Msgf("Error saving address 'City=%v;Street=%v;Tel=%v' from row=%d: %v", cty, str, t, rowCounter, addrerr)
				if ad == nil {
					logger.Fatal().Msgf("Error saving address 'City=%v;Street=%v;Tel=%v' from row=%d: %v", cty, str, t, rowCounter, addrerr)
				}
				businessCreate.AddAddresses(ad)
			} else {
				businessCreate.AddAddresses(addr)
			}
			_, err := businessCreate.Save(ctx)
			if err != nil {
				logger.Fatal().Msgf("Error saving business '%s' from row=%d: %v", record[0], rowCounter, err)
			}
			successfulImports++
		}
	}
	logger.Info().Msgf("Finished importing file '%s' with size %d bytes in %v - %d rows imported, %d rows skipped", stats.Name(), stats.Size(), time.Since(startTime), successfulImports, skippedImports)

}

func createBusinessFieldsandCheckAgainstTableHeader(header []string) ([]string, error) {
	var headersFromSchemas []string
	headersFromSchemas = utils.GetBusinessFields(headersFromSchemas, true, false)
	headersFromSchemas = utils.GetAddressFields(headersFromSchemas, true, false)

	if len(header) != len(headersFromSchemas) {
		return nil, fmt.Errorf("csv header length '%d' does not match needed fields in schema length '%d'", len(header), len(headersFromSchemas))
	}
	for i, headerField := range header {
		helper := strings.Split(headersFromSchemas[i], "$") // cut SchemaName
		if strings.ToUpper(headerField) != strings.ToUpper(helper[1]) {
			return nil, fmt.Errorf("header field '%s' does not match schema field '%s' (schema=%s)", headerField, helper[1], helper[0])
		}
	}
	return headersFromSchemas, nil
}

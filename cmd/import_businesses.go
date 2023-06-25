package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/business"
	_ "hynie.de/ohmab/ent/runtime"
	"hynie.de/ohmab/internal/pkg/config"
	"hynie.de/ohmab/internal/pkg/db"
	"hynie.de/ohmab/internal/pkg/log"
	"hynie.de/ohmab/internal/pkg/utils"
	"os"
	"strings"
	"time"
)

func main() {
	// Get logger
	logger := log.GetLoggerInstance()
	// Get the configuration
	configurations, err := config.GetConfiguration()
	if err != nil {
		logger.Fatal().Msgf("Error reading configurations: %v", err)
	}
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
	client, clientError := db.CreateClient(ctx, configurations)
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
		// Look into database first
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
					err := addressCreate.Mutation().SetField(schemaField, field)
					if err != nil {
						logger.Fatal().Msgf("Error setting address field '%s' to '%s': %v", schemaField, field, err)
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
				schemaHeader := headersFromSchemas[i]
				schema := strings.Split(schemaHeader, "$")[0]
				schemaField := strings.Split(schemaHeader, "$")[1]
				switch schema {
				case "BUSINESS":
					err := businessCreate.Mutation().SetField(schemaField, field)
					if err != nil {
						logger.Fatal().Msgf("Error setting business field '%s' to '%s': %v", schemaField, field, err)
					}
				case "ADDRESS":
					err := addressCreate.Mutation().SetField(schemaField, field)
					if err != nil {
						logger.Fatal().Msgf("Error setting address field '%s' to '%s' from row=%d: %v", schemaField, field, rowCounter, err)
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
				logger.Fatal().Msgf("Error saving address 'City=%v;Street=%v;Tel=%v' from row=%d: %v", cty, str, t, rowCounter, addrerr)
			}
			businessCreate.AddAddresses(addr)
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

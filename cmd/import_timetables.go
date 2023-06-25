package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"hynie.de/ohmab/ent"
	"hynie.de/ohmab/ent/address"
	"hynie.de/ohmab/ent/business"
	_ "hynie.de/ohmab/ent/runtime"
	schemas "hynie.de/ohmab/ent/schema"
	"hynie.de/ohmab/ent/schema/constants"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/internal/pkg/config"
	"hynie.de/ohmab/internal/pkg/db"
	"hynie.de/ohmab/internal/pkg/log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	dateTimeParseLayoutDots            = "02.01.2006 15:04"
	dateTimeParseLayoutDotsWithSeconds = "02.01.2006 15:04:05"
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
	logger.Debug().Msgf("Starting Timetables Import")
	filename := flag.String("filename", "import.csv", "Filename of the CSV file to import")
	allTimeentries := flag.Bool("alltimeentries", false, "Import also time entries in the past")
	flag.Parse()
	file, err := os.OpenFile(*filename, os.O_RDONLY, 0644)
	if err != nil {
		logger.Fatal().Msgf("Error opening file: %v", err)
	}
	stats, err := file.Stat()
	if err != nil {
		logger.Fatal().Msgf("Error getting stats: %v", err)
	}
	// CreateClient client
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
	for {
		record, readErr := reader.Read()
		if header {
			header = false
			headersFromSchemas_, err := createTimetableFieldsandCheckAgainstTableHeader(record)
			if err != nil {
				logger.Fatal().Msgf("Error checking table header: %v", err)
			} else {
				headersFromSchemas = headersFromSchemas_
				logger.Debug().Msgf("CSV headers successful validated")
			}
			continue
		}
		if record == nil {
			break
		}
		if readErr != nil {
			logger.Error().Msgf("Error while reading csv", readErr)
			break
		}
		// Get business and his address from db
		// Look into database first
		busQuery := client.Business.Query()
		var city, street, zip string
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
				switch strings.ToLower(schemaField) {
				case "zip":
					zip = field
				case "city":
					city = field
				case "street":
					street = field
				}
			case "TIMETABLE":
				continue
			default:
				logger.Fatal().Msgf("Unknown schema '%s'", schema)
			}

		}
		// get the business for this csv column first
		business_ := busQuery.Where(
			business.HasAddressesWith(address.Zip(zip), address.City(city), address.Street(street))).
			WithAddresses().Unique(true).FirstX(ctx)
		if business_ == nil { // businessAddresses_ cannot be nil
			skippedImports++
			logger.Info().Msgf("Business with address not found, Name/Zip/City/Street: %s/%s/%s/%s", record[0], zip, city, street)
			continue
		}

		// get the correct address
		businessAddresses_ := business_.Edges.Addresses
		var businessAddress *ent.Address
		for _, businessAddress_ := range businessAddresses_ {
			if businessAddress_.Zip == zip &&
				businessAddress_.City == city &&
				businessAddress_.Street == street {
				businessAddress = businessAddress_
			}
		}
		logger.Debug().Msgf("Processing address %v", businessAddress.ID)
		// get field types for Timetable
		tt := schemas.Timetable{}
		ttFields := tt.Fields()
		timetableCreate := client.Timetable.Create()
		for i, field := range record {
			if field == "" {
				continue
			}
			schemaHeader := headersFromSchemas[i]
			schema := strings.Split(schemaHeader, "$")[0]
			schemaField := strings.ToLower(strings.Split(schemaHeader, "$")[1])
			switch schema {
			case "BUSINESS":
				continue
			case "ADDRESS":
				continue
			case "TIMETABLE":
				if schemaField == "type" {
					switch strings.ToUpper(field) {
					case timetable.TypeEMERGENCYSERVICE.String():
						timetableCreate.Mutation().SetType(timetable.TypeEMERGENCYSERVICE)
					case timetable.TypeCLOSED.String():
						timetableCreate.Mutation().SetType(timetable.TypeCLOSED)
					case timetable.TypeHOLIDAY.String():
						timetableCreate.Mutation().SetType(timetable.TypeHOLIDAY)
					case timetable.TypeSPECIAL.String():
						timetableCreate.Mutation().SetType(timetable.TypeSPECIAL)
					case timetable.TypeREGULAR.String():
						timetableCreate.Mutation().SetType(timetable.TypeREGULAR)
					default:
						timetableCreate.Mutation().SetType(timetable.TypeDEFAULT)

					}
				} else {
					var fieldType string
					for _, ttField := range ttFields {
						ttFieldDesc := ttField.Descriptor()
						if ttFieldDesc.Name == schemaField {
							fieldType_ := ttFieldDesc.Info.Type.String()
							fieldType = fieldType_
							break
						}
					}
					var err_ error
					switch fieldType {
					case "time.Time":
						var dateOrTime time.Time
						if strings.HasPrefix(schemaField, "time") {
							time_, errParse := time.Parse(time.TimeOnly, field)
							if errParse != nil {
								logger.Fatal().Msgf("Error parsing time '%s': %v", field, errParse)
							}
							dateOrTime = time_.Local()
						} else {
							var l string
							// determine date input, if it contains dots, it is european format
							if strings.Contains(field, ".") {
								var l_ string
								if strings.Count(field, ":") > 1 {
									l_ = dateTimeParseLayoutDotsWithSeconds
								} else {
									l_ = dateTimeParseLayoutDots
								}
								l = l_
							} else {
								pattern_ := time.DateTime
								l = pattern_
							}
							date_, errParse := time.ParseInLocation(l, field, time.Local)
							if errParse != nil {
								logger.Fatal().Msgf("Error parsing date from value '%s': %v", field, errParse)
							}
							dateOrTime = date_
						}
						errSetter := timetableCreate.Mutation().SetField(schemaField, dateOrTime)
						err_ = errSetter
					case "bool":
						boolField, _ := strconv.ParseBool(field)
						errSetter := timetableCreate.Mutation().SetField(schemaField, boolField)
						err_ = errSetter
					default: // string
						errSetter := timetableCreate.Mutation().SetField(schemaField, field)
						err_ = errSetter
					}
					if err_ != nil {
						logger.Fatal().Msgf("Error setting timetable field '%s' to '%s': %v", schemaField, field, err)
					}
				}
			default:
				logger.Fatal().Msgf("Unknown schema '%s'", schema)
			}

		}
		var dateFrom, dateTo time.Time
		dateFrom, _ = timetableCreate.Mutation().DatetimeFrom()
		dateTo, _ = timetableCreate.Mutation().DatetimeTo()
		ttType, _ := timetableCreate.Mutation().GetType()
		// check if entry is in the past
		if !*allTimeentries && dateFrom.Before(time.Now()) {
			logger.Debug().Msgf("Timetable with type '%v' dateFrom %v is in the past, skipping", ttType, dateFrom)
			skippedImports++
			continue
		}
		exists := businessAddress.QueryTimetables().
			Where(timetable.DatetimeFrom(dateFrom)).
			Where(timetable.DatetimeTo(dateTo)).
			Where(timetable.TypeEQ(ttType)).
			CountX(ctx)
		if exists == 0 {
			timetableCreate.SetAddress(businessAddress).SaveX(ctx)
			successfulImports++
		} else {
			logger.Debug().Msgf("Timetable with type '%v' dateFrom %v and dateTo %v already exists", ttType, dateFrom, dateTo)
			skippedImports++
		}

	}
	logger.Info().Msgf("Finished importing file '%s' with size %d bytes in %v - %d rows imported, %d rows skipped", stats.Name(), stats.Size(), time.Since(startTime), successfulImports, skippedImports)

}

func createTimetableFieldsandCheckAgainstTableHeader(header []string) ([]string, error) {
	// unique id for a business and his address for a timetable entry - must be the first three columns in the csvSü
	headersFromSchemas := []string{
		"BUSINESS$NAME1",
		"ADDRESS$ZIP",
		"ADDRESS$CITY",
		"ADDRESS$STREET",
	}
	for _, field := range timetable.Columns {
		if strings.HasSuffix(field, constants.IDFieldName) {
			continue
		}
		if field == constants.TimeMixinCreatedAtFieldName ||
			field == constants.TimeMixinUpdatedAtFieldName ||
			field == constants.TimeMixinDeletedAtFieldName {
			continue
		}
		headersFromSchemas = append(headersFromSchemas, "TIMETABLE$"+field)
	}

	if len(header) != len(headersFromSchemas) {
		return nil, fmt.Errorf("csv header length '%d' does not match needed fields in schema length '%d', (perhaps wrong field limiter, should be ;), loaded header: %v", len(header), len(headersFromSchemas), header)
	}
	for i, headerField := range header {
		/*
			if i == 0 { // First column has to be NAME (for Business)
				if strings.ToUpper(headerField) != "NAME" {
					return nil, fmt.Errorf("first header field has to be NAME")
				}
				continue
			} else if i == 1 { // Second column has to be ZIP (for Business address)
				if strings.ToUpper(headerField) != "ZIP" {
					return nil, fmt.Errorf("second header field has to be ZIP")
				}
				continue
			} else if i == 2 { // Third column has to be CITY (for Business address)
				if strings.ToUpper(headerField) != "CITY" {
					return nil, fmt.Errorf("third header field has to be CITY")
				}
				continue
			}
		*/
		helper := strings.Split(headersFromSchemas[i], "$") // cut SchemaName
		if strings.ToUpper(headerField) != strings.ToUpper(helper[1]) {
			return nil, fmt.Errorf("header field '%s' does not match schema field '%s' (schema=%s)", headerField, helper[1], helper[0])
		}
	}
	return headersFromSchemas, nil
}

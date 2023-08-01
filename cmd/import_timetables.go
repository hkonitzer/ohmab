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
	"github.com/hkonitzer/ohmab/ent/schema/constants"
	"github.com/hkonitzer/ohmab/ent/timetable"
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
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
	config.Init()
	// Parse cmd line
	logger.Debug().Msgf("Starting Timetables Import")
	filename := flag.String("filename", "import.csv", "Filename of the CSV file to import")
	drop := flag.Bool("drop", false, "Drop all existing timetables before import")
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
	// Authorize me
	uv := privacy.UserViewer{Role: privacy.Admin}
	uv.SetUserID("import")
	ctx = privacy.NewContext(ctx, &uv)
	client, clientError := db.CreateClient(ctx)
	if clientError != nil {
		logger.Fatal().Msgf("Error creating client: %v", clientError)
	}
	defer client.Close()

	// Drop?
	if *drop {
		logger.Info().Msgf("Dropping all existing timetables")
		co, err := client.Timetable.Delete().Exec(ctx)
		if err != nil {
			logger.Fatal().Msgf("Error dropping timetables: %v", err)
		} else {
			logger.Debug().Msgf("Dropped %d timetables", co)
		}
	}

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
			schemaHeader := headersFromSchemas[i]
			schema := strings.Split(schemaHeader, "$")[0]
			schemaField := strings.Split(schemaHeader, "$")[1]
			switch schema {
			case "BUSINESS":
				switch strings.ToLower(schemaField) {
				case "alias":
					busQuery.Where(business.Alias(strings.ToUpper(field)))
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
		// if no city, street and zip is given, get the primary address
		if city == "" && street == "" && zip == "" {
			busQuery.Where().
				WithAddresses(
					func(q *ent.AddressQuery) {
						q.Where(address.PrimaryEQ(true)).FirstX(ctx)
					},
				).Unique(true)
		} else {
			busQuery.Where(
				business.HasAddressesWith(address.Zip(zip), address.City(city), address.Street(street))).
				WithAddresses().Unique(true)
		}
		business_ := busQuery.FirstX(ctx)
		if business_ == nil { // businessAddresses_ cannot be nil
			skippedImports++
			logger.Info().Msgf("Business with address not found, Name/Zip/City/Street: %s/%s/%s/%s", record[0], zip, city, street)
			continue
		}

		// get the correct address, only one give because of FirstX
		businessAddress := business_.Edges.Addresses[0]
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
				if schemaField == "timetable_type" {
					switch strings.ToUpper(field) {
					case timetable.TimetableTypeEMERGENCYSERVICE.String():
						timetableCreate.Mutation().SetTimetableType(timetable.TimetableTypeEMERGENCYSERVICE)
					case timetable.TimetableTypeCLOSED.String():
						timetableCreate.Mutation().SetTimetableType(timetable.TimetableTypeCLOSED)
					case timetable.TimetableTypeHOLIDAY.String():
						timetableCreate.Mutation().SetTimetableType(timetable.TimetableTypeHOLIDAY)
					case timetable.TimetableTypeSPECIAL.String():
						timetableCreate.Mutation().SetTimetableType(timetable.TimetableTypeSPECIAL)
					case timetable.TimetableTypeREGULAR.String():
						timetableCreate.Mutation().SetTimetableType(timetable.TimetableTypeREGULAR)
					default:
						timetableCreate.Mutation().SetTimetableType(timetable.TimetableTypeDEFAULT)
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
						err_ = timetableCreate.Mutation().SetField(schemaField, dateOrTime)
					case "bool":
						boolField, _ := strconv.ParseBool(field)
						err_ = timetableCreate.Mutation().SetField(schemaField, boolField)
						if err_ != nil {
							logger.Fatal().Msgf("Error setting timetable field '%s' to '%s': %v", schemaField, field, err_)
						}
					case "uint8":
						uint64Value, _ := strconv.ParseUint(field, 10, 8)
						uint8Value := uint8(uint64Value)
						err_ = timetableCreate.Mutation().SetField(schemaField, uint8Value)
						if err_ != nil {
							logger.Fatal().Msgf("Error setting timetable field '%s' to '%s': %v", schemaField, field, err_)
						}
					case "int":
						i, _ := strconv.Atoi(field)
						err_ = timetableCreate.Mutation().SetField(schemaField, i)
					default: // string
						err_ = timetableCreate.Mutation().SetField(schemaField, field)
						if err_ != nil {
							logger.Fatal().Msgf("Error setting timetable field '%s' to '%s': %v", schemaField, field, err_)
						}
					}
				}
			default:
				logger.Fatal().Msgf("Unknown schema '%s'", schema)
			}

		}
		// validate duration against dateTo

		var dateFrom, dateTo time.Time
		dateFrom, _ = timetableCreate.Mutation().DatetimeFrom()
		dateTo, _ = timetableCreate.Mutation().DatetimeTo()
		ttType, _ := timetableCreate.Mutation().TimetableType()
		// check if entry is in the past
		if !*allTimeentries && dateFrom.Before(time.Now()) {
			logger.Debug().Msgf("Timetable with type '%v' dateFrom %v is in the past, skipping", ttType, dateFrom)
			skippedImports++
			continue
		}
		exists := businessAddress.QueryTimetables().
			Where(timetable.DatetimeFrom(dateFrom)).
			Where(timetable.DatetimeTo(dateTo)).
			Where(timetable.TimetableTypeEQ(ttType)).
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
		"BUSINESS$ALIAS",
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
		helper := strings.Split(headersFromSchemas[i], "$") // cut SchemaName
		if strings.ToUpper(headerField) != strings.ToUpper(helper[1]) {
			return nil, fmt.Errorf("header field '%s' does not match schema field '%s' (schema=%s)", headerField, helper[1], helper[0])
		}
	}
	return headersFromSchemas, nil
}

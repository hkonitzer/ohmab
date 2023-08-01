package main

import (
	"context"
	"encoding/csv"
	"github.com/hkonitzer/ohmab/ent"
	_ "github.com/hkonitzer/ohmab/ent/runtime"
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/hkonitzer/ohmab/internal/pkg/utils"
	"os"
	"strconv"
)

var exportlog = log.GetLoggerInstance()
var ctx = context.Background()

func exportBusinesses(client *ent.Client) {
	exportlog.Debug().Msgf("Starting Businesses Export")
	// create output file
	file, err := os.Create("./businesses.csv")
	if err != nil {
		exportlog.Fatal().Msgf("Error creating file: %v", err)
	}
	// export businesses

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	writer.UseCRLF = true

	defer writer.Flush()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var headers []string
	headers = utils.GetBusinessFields(headers, false, true)
	headers = utils.GetAddressFields(headers, false, true)
	wErr := writer.Write(headers)
	if wErr != nil {
		exportlog.Fatal().Msgf("Error writing headers: %v", wErr)
	}
	b, err := client.Business.Query().WithAddresses().All(ctx)
	if err != nil {
		exportlog.Fatal().Msgf("Error querying businesses: %v", err)
	}
	for _, business := range b {
		var bRow []string
		bRow = append(bRow, business.Name1)
		bRow = append(bRow, business.Name2)
		bRow = append(bRow, business.Alias)
		bRow = append(bRow, business.Telephone)
		bRow = append(bRow, business.Email)
		bRow = append(bRow, business.Website)
		bRow = append(bRow, business.Comment)
		for _, addr := range business.Edges.Addresses {
			aRow := bRow
			aRow = append(aRow, addr.Addition)
			aRow = append(aRow, addr.Street)
			aRow = append(aRow, addr.City)
			aRow = append(aRow, addr.Zip)
			aRow = append(aRow, addr.State)
			aRow = append(aRow, addr.Country)
			aRow = append(aRow, addr.Locale)
			aRow = append(aRow, strconv.FormatBool(addr.Primary))
			aRow = append(aRow, addr.Telephone)
			aRow = append(aRow, addr.Comment)
			err := writer.Write(aRow)
			if err != nil {
				exportlog.Fatal().Msgf("Error writing row: %v", err)
			}
		}

	}
	writer.Flush()
	exportlog.Debug().Msgf("Businesses Export finished with %d businesses", len(b))
}

func exportTimetables(client *ent.Client) {
	exportlog.Debug().Msgf("Starting Timetables Export")
	// create output file
	file, err := os.Create("./timetables.csv")
	if err != nil {
		exportlog.Fatal().Msgf("Error creating file: %v", err)
	}
	// export businesses

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	writer.UseCRLF = true

	defer writer.Flush()
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	headers := []string{
		"ALIAS", "ZIP", "CITY", "STREET",
		"TIMETABLE_TYPE", "DATETIME_FROM", "DURATION", "DATETIME_TO", "TIME_WHOLE_DAY", "COMMENT", "AVAILABILITY_BY_PHONE",
		"AVAILABILITY_BY_EMAIL", "AVAILABILITY_BY_SMS", "AVAILABILITY_BY_WHATSAPP",
	}
	wErr := writer.Write(headers)
	if wErr != nil {
		exportlog.Fatal().Msgf("Error writing headers: %v", wErr)
	}
	a, err := client.Address.Query().WithTimetables().WithBusiness().All(ctx)
	if err != nil {
		exportlog.Fatal().Msgf("Error querying businesses: %v", err)
	}
	c := 0
	for _, addr := range a {
		tt := addr.Edges.Timetables
		if len(tt) == 0 {
			continue
		}
		var row []string
		for _, t := range tt {
			row = append(row, addr.Edges.Business.Alias)
			row = append(row, addr.Zip)
			row = append(row, addr.City)
			row = append(row, addr.Street)
			row = append(row, t.TimetableType.String())
			row = append(row, t.DatetimeFrom.Format("02.01.2006 15:04"))
			row = append(row, strconv.Itoa(t.Duration))
			row = append(row, t.DatetimeTo.Format("02.01.2006 15:04"))
			row = append(row, strconv.FormatBool(t.TimeWholeDay))
			row = append(row, t.Comment)
			if t.AvailabilityByPhone == "" {
				row = append(row, t.AvailabilityByPhone)
			} else {
				row = append(row, "")
			}
			if t.AvailabilityByEmail == "" {
				row = append(row, t.AvailabilityByEmail)
			} else {
				row = append(row, "")
			}
			if t.AvailabilityBySms == "" {
				row = append(row, t.AvailabilityBySms)
			} else {
				row = append(row, "")
			}
			if t.AvailabilityByWhatsapp == "" {
				row = append(row, t.AvailabilityByWhatsapp)
			} else {
				row = append(row, "")
			}
			err := writer.Write(row)
			row = []string{}
			if err != nil {
				exportlog.Fatal().Msgf("Error writing row: %v", err)
			}
			c++
		}

	}
	writer.Flush()
	exportlog.Debug().Msgf("Timetables Export finished with %d timetables", c)
}

func main() {
	// Get the configuration
	config.Init()
	// Create client

	ctx = privacy.NewContext(ctx, &privacy.UserViewer{Role: privacy.Admin})
	client, clientError := db.CreateClient(ctx)
	if clientError != nil {
		exportlog.Fatal().Msgf("Error creating client: %v", clientError)
	}
	defer client.Close()

	exportBusinesses(client)
	exportTimetables(client)
}

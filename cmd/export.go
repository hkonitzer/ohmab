package main

import (
	"context"
	"encoding/csv"
	_ "github.com/hkonitzer/ohmab/ent/runtime"
	"github.com/hkonitzer/ohmab/internal/pkg/common/config"
	"github.com/hkonitzer/ohmab/internal/pkg/common/log"
	"github.com/hkonitzer/ohmab/internal/pkg/db"
	"github.com/hkonitzer/ohmab/internal/pkg/privacy"
	"github.com/hkonitzer/ohmab/internal/pkg/utils"
	"os"
	"strconv"
)

func main() {
	// Get logger
	logger := log.GetLoggerInstance()
	// Get the configuration
	config.Init()
	logger.Debug().Msgf("Starting Businesses Export")
	// create output file
	//file, err := os.OpenFile("./businesses.csv", os.O_CREATE, 0664)
	file, err := os.Create("./businesses.csv")
	if err != nil {
		logger.Fatal().Msgf("Error creating file: %v", err)
	}
	// Create client
	ctx := context.TODO()
	ctx = privacy.NewContext(ctx, &privacy.UserViewer{Role: privacy.Admin})
	client, clientError := db.CreateClient(ctx)
	if clientError != nil {
		logger.Fatal().Msgf("Error creating client: %v", clientError)
	}
	defer client.Close()
	// export businesses

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	writer.UseCRLF = true

	defer writer.Flush()
	defer file.Close()

	var headers []string
	headers = utils.GetBusinessFields(headers, false, true)
	headers = utils.GetAddressFields(headers, false, true)
	wErr := writer.Write(headers)
	if wErr != nil {
		logger.Fatal().Msgf("Error writing headers: %v", wErr)
	}
	b, err := client.Business.Query().WithAddresses().All(ctx)
	if err != nil {
		logger.Fatal().Msgf("Error querying businesses: %v", err)
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
				logger.Fatal().Msgf("Error writing row: %v", err)
			}
		}

	}
	writer.Flush()
}

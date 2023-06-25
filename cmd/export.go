package main

import (
	"context"
	"encoding/csv"
	"hynie.de/ohmab/internal/pkg/config"
	"hynie.de/ohmab/internal/pkg/db"
	"hynie.de/ohmab/internal/pkg/log"
	"hynie.de/ohmab/internal/pkg/utils"
	"os"
)

func main() {
	// Get logger
	logger := log.GetLoggerInstance()
	// Get the configuration
	configurations, err := config.GetConfiguration()
	if err != nil {
		logger.Fatal().Msgf("Error reading configurations: %v", err)
	}
	logger.Debug().Msgf("Starting Businesses Export")
	// create output file
	//file, err := os.OpenFile("./businesses.csv", os.O_CREATE, 0664)
	file, err := os.Create("./businesses.csv")
	if err != nil {
		logger.Fatal().Msgf("Error creating file: %v", err)
	}
	// Create client
	ctx := context.TODO()
	client, clientError := db.CreateClient(ctx, configurations)
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

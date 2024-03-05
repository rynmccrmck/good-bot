package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	goodbot "github.com/rynmccrmck/good-bot"
)

// BulkVerify reads an input CSV and writes the results to an output CSV.
func BulkVerify(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return processCSV(inputFile, outputFile)
}

func processCSV(inputFile io.Reader, outputFile io.Writer) error {
	reader := csv.NewReader(inputFile)
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Assuming the input CSV has headers and the first two columns are 'user_agent' and 'ip_address'
	headers, err := reader.Read()
	if err != nil {
		fmt.Printf("Error reading headers: %v\n", err)
		os.Exit(1)
	}

	headers = append(headers, "bot_status", "bot_name")
	writer.Write(headers)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading records: %v\n", err)
		os.Exit(1)
	}

	for _, record := range records {
		ua := record[0]
		ip := record[1]

		botResult, _ := goodbot.CheckBotStatus(ua, ip)

		var result string
		switch botResult.BotStatus {
		case goodbot.BotStatusFriendly:
			result = "friendly"
		case goodbot.BotStatusMalicious:
			result = "malicious"
		case goodbot.BotStatusUnknown:
			result = "unknown"
		case goodbot.BotStatusPotentialImposter:
			result = "potential_imposter"
		case goodbot.BotStatusPotentiallyFriendly:
			result = "potentially_friendly"
		}
		record = append(record, result, botResult.BotName)
		writer.Write(record)
	}

	return nil
}

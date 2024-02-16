package goodbot

import (
	"encoding/json"
	"fmt"
	"os"
)

// BotData represents the structure of our JSON data. Adjust this according to your actual JSON structure.
type BotData map[string]map[string]struct {
	UserAgent        string   `json:"User Agent"`
	Method           string   `json:"Method"`
	ValidDomains     []string `json:"Valid domains"`
	BypassFlag       string   `json:"Bypass flag,omitempty"`
	UserAgentPattern string   `json:"User Agent Pattern,omitempty"`
}

// LoadBotData loads bot data from a specified JSON file path.
func LoadBotData(filePath string) (BotData, error) {
	var botData BotData
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening bot data file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&botData); err != nil {
		return nil, fmt.Errorf("error decoding bot data: %v", err)
	}

	return botData, nil
}

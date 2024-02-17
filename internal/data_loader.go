package goodbot

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed data/goodbots.json
var goodbotsData embed.FS

var bots BotData

func init() {
	var err error
	bots, err = loadBotData()
	if err != nil {
		panic(err)
	}
}

type BotEntry struct {
	Name             string   `json:"Name"`
	Method           string   `json:"Method"`
	ValidDomains     []string `json:"ValidDomains"`
	UserAgentPattern string   `json:"UserAgentPattern,omitempty"`
}

type BotData struct {
	Bots []BotEntry `json:"Bots"`
}

func loadBotData() (BotData, error) {
	var bots BotData

	// ReadFile returns a byte slice that you can then unmarshal into your struct.
	data, err := goodbotsData.ReadFile("data/goodbots.json")
	if err != nil {
		return bots, fmt.Errorf("failed to read embedded data: %w", err)
	}

	if err := json.Unmarshal(data, &bots); err != nil {
		return bots, fmt.Errorf("failed to unmarshal bot data: %w", err)
	}

	return bots, nil
}

func GetBots() BotData {
	return bots
}

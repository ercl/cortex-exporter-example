package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Instruments []InstrumentConfig `json:"instrumentConfigs"`
}

type InstrumentConfig struct {
	Type           string `json:"type"`
	Label          string `json:"label"`
	Description    string `json:"description"`
	DataPointCount int    `json:"dataPointCount"`
	RecordInterval int    `json:"recordInterval"`
}

func readConfig(filepath string) (*Config, error) {
	// Read JSON file.
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config Config
	json.Unmarshal([]byte(file), &config)
	return &config, nil
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	Instruments []InstrumentConfig `json:"instrumentConfigs"`
}

type InstrumentConfig struct {
	Type           string        // Instrument type e.g counter
	Label          string        // Label for instrument e.g a.counter
	Description    string        // Description for the instrument
	DataPointCount int           // Number of values to record for the instrument
	RecordInterval time.Duration // Duration between recording a new data point
}

func readConfig(filepath string) error {
	// Read JSON file.
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	var config Config
	json.Unmarshal([]byte(file), &config)
	return nil
}

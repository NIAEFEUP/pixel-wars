package model

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration stores the initialize the canvas and other runtime variables.
type Configuration struct {
	CanvasWidth     int    `json:"canvasWidth"`
	CanvasHeight    int    `json:"canvasHeight"`
	PixelsPerMinute int    `json:"pixelsPerMinute"`
	DebugMode       bool   `json:"debugMode"`
	Host            string `json:"host"`
}

// LoadConfigurationFile loads the necessary configuration from a json file to setup redis if necessary
func LoadConfigurationFile() Configuration {
	dat, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("Err Couldn't read config file: %v Exiting...\n", err)
		os.Exit(1)
	}
	config := Configuration{}
	errJSON := json.Unmarshal(dat, &config)
	if errJSON != nil {
		fmt.Printf("Err Couldn't read config file: %v Exiting...\n", err)
		os.Exit(1)
	}
	return config
}

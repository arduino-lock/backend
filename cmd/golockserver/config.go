package main

import (
	"encoding/json"
	"os"
)

type config struct {
	Development bool
	Host        string
	Port        int
	CardsDB     string
}

func loadConfig(path string) (*config, error) {
	cfg := &config{}

	var cfgPath string
	if len(path) < 1 {
		cfgPath = "config.json"
	} else {
		cfgPath = path
	}

	f, err := os.Open(cfgPath)
	if err != nil {
		return cfg, err
	}

	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(&cfg)

	return cfg, err
}

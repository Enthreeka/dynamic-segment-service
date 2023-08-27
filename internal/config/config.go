package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		Postgres Postgres `json:"postgres"`
	}

	Postgres struct {
		URL string `json:"url"`
	}
)

func New(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

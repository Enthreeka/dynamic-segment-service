package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		Server   Server   `json:"server"`
		Postgres Postgres `json:"postgres"`
	}

	Server struct {
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"typeserver"`
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

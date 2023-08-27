package main

import (
	"github.com/Enthreeka/dynamic-segment-service/internal/config"
	"github.com/Enthreeka/dynamic-segment-service/internal/server"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
)

func main() {
	path := `configs/config.json`

	log := logger.New()

	cfg, err := config.New(path)
	if err != nil {
		log.Error("failed to load config: %v", err)
	}

	if err := server.Run(cfg, log); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}

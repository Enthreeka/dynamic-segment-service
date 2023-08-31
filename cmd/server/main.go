package main

import (
	_ "github.com/Enthreeka/dynamic-segment-service/docs"
	"github.com/Enthreeka/dynamic-segment-service/internal/config"
	"github.com/Enthreeka/dynamic-segment-service/internal/server"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger Api for Dynamic User Segmentation Service

// @host localhost:8080
// @BasePath /
func main() {
	path := `configs/config.json`

	log := logger.New()

	cfg, err := config.New(path)
	if err != nil {
		log.Error("failed to load config: %v", err)
	}

	//--------------------------------------------------------
	//date, err := time.Parse("2006-01-02", "2023-08-30")
	//if err != nil {
	//	log.Error("%v", err)
	//}
	//r := csv.Record{}
	//r.Read("0fa0c12b-32a5-49ef-bbb5-3539f8f67971", "AVITO_CAARRR_FFGD", date)
	//--------------------------------------------------------

	if err := server.Run(cfg, log); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}

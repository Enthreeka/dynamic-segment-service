package server

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/config"
	"github.com/Enthreeka/dynamic-segment-service/internal/repo"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
	"github.com/Enthreeka/dynamic-segment-service/pkg/postgres"
)

func Run(cfg *config.Config, log *logger.Logger) error {

	conn, err := postgres.New(context.Background(), cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}

	userRepo := repo.NewUserReposotory(conn)

	//user, err := userRepo.GetALL(context.Background())
	user, err := userRepo.GetByID(context.Background(), "70c247da-377a-42ac-97f6-316abfc43722")
	if err != nil {
		log.Error("failed to get segments - %v", err)
	}

	log.Info("%#v", user)

	return nil
}

package server

import (
	"context"
	"fmt"
	"github.com/Enthreeka/dynamic-segment-service/internal/config"
	controller "github.com/Enthreeka/dynamic-segment-service/internal/controller/http"
	"github.com/Enthreeka/dynamic-segment-service/internal/repo"
	"github.com/Enthreeka/dynamic-segment-service/internal/usecase"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
	"github.com/Enthreeka/dynamic-segment-service/pkg/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Run(cfg *config.Config, log *logger.Logger) error {

	conn, err := postgres.New(context.Background(), cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}

	defer conn.Close()

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	segmentRepo := repo.NewSegmentRepository(conn)
	userRepo := repo.NewUserReposotory(conn)

	segmentUsecase := usecase.NewSegmentService(segmentRepo, log)
	userUsecase := usecase.NewUserService(userRepo, log)

	segmentHandler := controller.NewSegmentHandler(segmentUsecase, log)
	userHandler := controller.NewUserHandler(segmentUsecase, userUsecase, log)

	api := app.Group("/api")

	v1 := api.Group("/segment")
	v1.Get("/", segmentHandler.GetAll)
	v1.Post("/", segmentHandler.CreateSegment)
	v1.Delete("/:segment", segmentHandler.DeleteSegment)

	v2 := api.Group("/user")
	v2.Post("/:id", userHandler.GetUserSegment)
	v2.Post("/", userHandler.SetSegments)
	v2.Delete("/:segment", userHandler.DeleteSegments)
	v2.Get("/all", userHandler.GetAllUser)

	log.Info("Starting http server: %s:%s", cfg.Server.TypeServer, cfg.Server.Port)

	if err = app.Listen(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}

package http

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/apperror"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/internal/usecase"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	segmentUsecase usecase.SegmentService
	userUsecase    usecase.UserService

	log *logger.Logger
}

func NewUserHandler(segmentUsecase usecase.SegmentService, userUsecase usecase.UserService, log *logger.Logger) *userHandler {
	return &userHandler{
		segmentUsecase: segmentUsecase,
		userUsecase:    userUsecase,
		log:            log,
	}
}

func (u *userHandler) GetUserSegment(c *fiber.Ctx) error {
	u.log.Info("getting user segments started")

	user := new(entity.User)
	err := c.BodyParser(&user)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
			"user_id": user.ID,
		})
	}

	userInfo, err := u.userUsecase.GetUserInfo(context.Background(), user.ID)
	if err != nil {
		u.log.Error("failed to get info about user - %s: %v", user.ID, err)

		if err == apperror.ErrSegmentsNotFound {
			return c.Status(fiber.StatusNotFound).JSON(apperror.ErrSegmentsNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	u.log.Info("getting user segments completed successfully")
	userInfo.ID = user.ID
	return c.Status(fiber.StatusOK).JSON(userInfo)
}

func (u *userHandler) SetSegments(c *fiber.Ctx) error {
	u.log.Info("setting segments to user started")

	user := new(entity.User)
	err := c.BodyParser(user)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	for indx := range user.Segments {
		segmentID, err := u.segmentUsecase.GetIDByName(context.Background(), user.Segments[indx].Segment)
		if err != nil {
			u.log.Error("failed to get id segment by name")

			if err == apperror.ErrSegmentsNotFound {
				return c.Status(fiber.StatusNotFound).JSON(apperror.ErrSegmentsNotFound)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		user.Segments[indx].ID = segmentID.ID
	}

	err = u.userUsecase.SetSegment(context.Background(), user)
	if err != nil {
		u.log.Error("failed to set segments to user - %s: %v", user.ID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	u.log.Info("setting segments to user completed successfully")
	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"message":        "Completed successfully",
		"added_segments": user.Segments,
	})
}

func (u *userHandler) DeleteSegments(c *fiber.Ctx) error {
	u.log.Info("deleting segments from user started")

	user := new(entity.User)
	err := c.BodyParser(user)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	for indx := range user.Segments {
		segmentID, err := u.segmentUsecase.GetIDByName(context.Background(), user.Segments[indx].Segment)
		if err != nil {
			u.log.Error("failed to get id segment by name")

			if err == apperror.ErrSegmentsNotFound {
				return c.Status(fiber.StatusNotFound).JSON(apperror.ErrSegmentsNotFound)
			}
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		user.Segments[indx].ID = segmentID.ID
	}

	err = u.userUsecase.DeleteUserSegment(context.Background(), user)
	if err != nil {
		u.log.Error("failed to delete segments from user - %s: %v", user.ID, err)

		if err == apperror.ErrSegmentsNotFound {
			return c.Status(fiber.StatusNotFound).JSON(apperror.ErrSegmentsNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	u.log.Info("deleting segments from user completed successfully")
	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"message":          "Completed successfully",
		"deleted_segments": user.Segments,
	})
}

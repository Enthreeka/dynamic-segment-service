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

// swagger:parameters UUID
type UUID struct {
	UserID string `json:"id"`
}

// GetUserSegment godoc
// @Summary Get User Segments
// @Tags user
// @Description get segments associated with a user
// @Accept json
// @Produce json
// @Param input body entity.UUID true "User ID"
// @Success 200 {object} entity.User
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/user/:id [post]
func (u *userHandler) GetUserSegment(c *fiber.Ctx) error {
	u.log.Info("getting user segments started")

	var input UUID
	err := c.BodyParser(&input)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
			"user_id": input.UserID,
		})
	}

	userInfo, err := u.userUsecase.GetUserInfo(context.Background(), input.UserID)
	if err != nil {
		u.log.Error("failed to get info about user - %s: %v", input.UserID, err)

		if err == apperror.ErrUserNotFound {
			return c.Status(fiber.StatusNotFound).JSON(apperror.ErrUserNotFound)
		}
		if err == apperror.ErrSegmentsNotFound {
			return c.Status(fiber.StatusNotFound).JSON(apperror.ErrSegmentsNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	userInfo.ID = input.UserID
	u.log.Info("getting user segments completed successfully")
	return c.Status(fiber.StatusOK).JSON(userInfo)
}

// SetSegments godoc
// @Summary Set User Segments
// @Tags user
// @Description set segments for a user
// @Accept json
// @Produce json
// @Param input body entity.User true "User segments to set"
// @Success 200 {object} map[string]interface{}
// @Failure 302 {object} apperror.AppError
// @Failure 400,404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/user [post]
func (u *userHandler) SetSegments(c *fiber.Ctx) error {
	u.log.Info("setting segments to user started")

	user := new(entity.User)
	err := c.BodyParser(user)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewAppError(err, "Invalid request body"))
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

		if err == apperror.ErrUserHasSegment {
			return c.Status(fiber.StatusFound).JSON(apperror.ErrUserHasSegment)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	u.log.Info("setting segments to user completed successfully")
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message":        "Completed successfully",
		"added_segments": user.Segments,
	})
}

// DeleteSegments godoc
// @Summary Delete User Segments
// @Tags user
// @Description delete segments from a user
// @Accept json
// @Produce json
// @Param input body entity.User true "User segments to delete"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/user/:segments [delete]
func (u *userHandler) DeleteSegments(c *fiber.Ctx) error {
	u.log.Info("deleting segments from user started")

	user := new(entity.User)
	err := c.BodyParser(user)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewAppError(err, "Invalid request body"))
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
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message":          "Completed successfully",
		"deleted_segments": user.Segments,
	})
}

// GetAllUser godoc
// @Summary Get All User
// @Tags user
// @Description get all users and their segments
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]string
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/user/all [get]
func (u *userHandler) GetAllUser(c *fiber.Ctx) error {
	u.log.Info("getting all users started")

	users, err := u.userUsecase.GetAllUser(context.Background())
	if err != nil {
		if err == apperror.ErrUsersNotFound {
			u.log.Error("failed to ger all user: %v", err)
			return c.Status(fiber.StatusNotFound).JSON(apperror.ErrUsersNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

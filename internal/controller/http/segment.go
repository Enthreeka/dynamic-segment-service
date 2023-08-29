package http

import (
	"context"
	_ "github.com/Enthreeka/dynamic-segment-service/docs"
	"github.com/Enthreeka/dynamic-segment-service/internal/apperror"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/internal/usecase"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
	"github.com/Enthreeka/dynamic-segment-service/pkg/validation"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type segmentHandler struct {
	segmentUsecase usecase.SegmentService
	log            *logger.Logger
}

func NewSegmentHandler(segmentUsecase usecase.SegmentService, log *logger.Logger) *segmentHandler {
	return &segmentHandler{
		segmentUsecase: segmentUsecase,
		log:            log,
	}
}

// CreateSegment godoc
// @Summary Create Segment
// @Tags segment
// @Description create segment
// @Accept json
// @Produce json
// @Param input body entity.Segment true "Segment data to be created"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/segment [post]
func (s *segmentHandler) CreateSegment(c *fiber.Ctx) error {
	s.log.Info("start of segment creation")

	segment := new(entity.Segment)
	err := c.BodyParser(&segment)
	if err != nil {
		s.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewAppError(err, "Invalid request body"))
	}

	segment.Segment = strings.ToUpper(segment.Segment)
	validSegment, err := validation.ValidSegmentName(segment.Segment)
	if !validSegment {
		s.log.Error("invalid data: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewAppError(err, "invalid input data"))
	}

	_, err = s.segmentUsecase.GetIDByName(context.Background(), segment.Segment)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewAppError(err, "segment already exist"))
	}

	err = s.segmentUsecase.CreateSegment(context.Background(), segment.Segment)
	if err != nil {
		s.log.Error("failed to create segment: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	s.log.Info("segment creation completed successfully : %s", segment.Segment)
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message":         "Completed successfully",
		"created_segment": segment.Segment,
	})
}

// DeleteSegment godoc
// @Summary Delete Segment
// @Tags segment
// @Description delete segment
// @Accept json
// @Produce json
// @Param input body entity.Segment true "Segment data to be deleted"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/segment/:segment [delete]
func (s *segmentHandler) DeleteSegment(c *fiber.Ctx) error {
	s.log.Info("start of segment deletion")

	segment := new(entity.Segment)
	err := c.BodyParser(&segment)
	if err != nil {
		s.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	searchSegment, err := s.segmentUsecase.GetIDByName(context.Background(), segment.Segment)
	if err != nil {
		s.log.Error("failed to get segment: %v", err)

		if err == apperror.ErrSegmentsNotFound {
			c.Status(fiber.StatusNotFound).JSON(apperror.ErrSegmentsNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	err = s.segmentUsecase.DeleteSegmentByName(context.Background(), searchSegment)
	if err != nil {
		s.log.Error("failed to delete segment: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	s.log.Info("segment deletion completed successfully : %s", segment.Segment)
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message":         "Completed successfully",
		"deleted_segment": segment.Segment,
	})
}

// GetAll godoc
// @Summary Get All Segments
// @Tags segment
// @Description get all segments
// @Accept  json
// @Produce  json
// @Success 200 {object} map[int]string "segmentsMap"
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/segment [get]
func (s *segmentHandler) GetAll(c *fiber.Ctx) error {
	s.log.Info("start of searching segments")

	segments, err := s.segmentUsecase.GetAllSegments(context.Background())
	if err != nil {
		s.log.Error("failed to get all segments: %v", err)

		if err == apperror.ErrSegmentsNotFound {
			return c.Status(fiber.StatusNotFound).JSON(apperror.ErrSegmentsNotFound)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	segmentsMap := make(map[int]string)
	for _, el := range segments {
		segmentsMap[el.ID] = el.Segment
	}

	s.log.Info("search segments completed successfully")
	return c.Status(fiber.StatusOK).JSON(segmentsMap)
}

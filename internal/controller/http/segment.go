package http

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/internal/usecase"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
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

func (s *segmentHandler) CreateSegment(c *fiber.Ctx) error {
	s.log.Info("start of segment creation")

	segment := new(entity.Segment)
	err := c.BodyParser(&segment)
	if err != nil {
		s.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	//TODO Нужно сделать apperror как в notes service
	err = s.segmentUsecase.CreateSegment(context.Background(), segment.Segment)
	if err != nil {
		s.log.Error("failed to create segment: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err,
		})
	}

	s.log.Info("segment creation completed successfully : %s", segment.Segment)
	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"message":         "Completed successfully",
		"created_segment": segment.Segment,
	})
}

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

	//TODO Нужно сделать apperror как в notes service
	searchSegment, err := s.segmentUsecase.GetIDByName(context.Background(), segment.Segment)
	if err != nil {
		s.log.Error("failed to get segment: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Надо сделать ошибку что нет такого сегмента",
		})
	}

	//TODO Нужно сделать apperror как в notes service
	err = s.segmentUsecase.DeleteSegmentByName(context.Background(), searchSegment)
	if err != nil {
		s.log.Error("failed to delete segment: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err,
		})
	}

	s.log.Info("segment deletion completed successfully : %s", segment.Segment)
	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"message":         "Completed successfully",
		"deleted_segment": segment.Segment,
	})
}

func (s *segmentHandler) GetAll(c *fiber.Ctx) error {
	s.log.Info("start of searching segments")

	//segment := make([]entity.Segment, 0)
	//err := c.BodyParser(&segment)
	//if err != nil {
	//	s.log.Error("failed to parse request body: %v", err)
	//	return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
	//		"message": "Invalid request body",
	//	})
	//}

	//TODO Нужно сделать apperror как в notes service
	segments, err := s.segmentUsecase.GetAllSegments(context.Background())
	if err != nil {
		s.log.Error("failed to get all segments: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Надо сделать ошибку ",
		})
	}

	segmentsMap := make(map[int]string)
	for _, el := range segments {
		segmentsMap[el.ID] = el.Segment
	}

	s.log.Info("search segments completed successfully")
	return c.Status(fiber.StatusCreated).JSON(segmentsMap)

}

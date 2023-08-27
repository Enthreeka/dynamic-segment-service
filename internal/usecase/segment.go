package usecase

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/internal/repo"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
	"github.com/Enthreeka/dynamic-segment-service/pkg/validation"
)

type segmentService struct {
	segmentRepo repo.SegmentRepository
	log         *logger.Logger
}

func NewSegmentService(segmentRepo repo.SegmentRepository, log *logger.Logger) *segmentService {
	return &segmentService{
		segmentRepo: segmentRepo,
		log:         log,
	}
}

func (s *segmentService) CreateSegment(ctx context.Context, segment string) error {
	validSegment, err := validation.ValidSegmentName(segment)
	if !validSegment {
		return err
	}

	err = s.segmentRepo.Create(ctx, segment)
	if err != nil {
		return err
	}

	return nil
}

func (s *segmentService) DeleteSegmentByName(ctx context.Context, segment *entity.Segment) error {
	err := s.segmentRepo.DeleteByID(ctx, segment)
	if err != nil {
		return err
	}

	return nil
}

func (s *segmentService) GetIDByName(ctx context.Context, segmentType string) (*entity.Segment, error) {
	segment, err := s.segmentRepo.GetByName(ctx, segmentType)
	if err != nil {
		return nil, err
	}

	return segment, nil
}

func (s *segmentService) GetAllSegments(ctx context.Context) ([]entity.Segment, error) {
	segments, err := s.segmentRepo.GetALL(ctx)
	if err != nil {
		return nil, err
	}

	return segments, nil
}

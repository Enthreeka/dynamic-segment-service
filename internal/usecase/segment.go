package usecase

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/apperror"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/internal/repo"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
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
	//validSegment, err := validation.ValidSegmentName(segment)
	//if !validSegment {
	//	return apperror.NewAppError(err, "invalid input data")
	//}

	err := s.segmentRepo.Create(ctx, segment)
	if err != nil {
		return apperror.NewAppError(err, "failed to create segment")
	}

	return nil
}

func (s *segmentService) DeleteSegmentByName(ctx context.Context, segment *entity.Segment) error {
	err := s.segmentRepo.DeleteByID(ctx, segment)
	if err != nil {
		return apperror.NewAppError(err, "failed to delete segment")
	}

	return nil
}

func (s *segmentService) GetIDByName(ctx context.Context, segmentType string) (*entity.Segment, error) {
	segment, err := s.segmentRepo.GetByName(ctx, segmentType)
	if err != nil {
		if err == apperror.ErrSegmentsNotFound {
			return nil, apperror.ErrSegmentsNotFound
		}
		return nil, apperror.NewAppError(err, "failed to get id segment by name")
	}
	return segment, nil
}

func (s *segmentService) GetAllSegments(ctx context.Context) ([]entity.Segment, error) {
	segments, err := s.segmentRepo.GetALL(ctx)
	if err != nil {
		return nil, apperror.NewAppError(err, "failed to get all segments")
	}

	if len(segments) == 0 {
		return nil, apperror.ErrSegmentsNotFound
	}

	return segments, nil
}

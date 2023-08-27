package repo

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/pkg/postgres"
)

type segmentRepository struct {
	*postgres.Postgres
}

func NewSegmentRepository(postgres *postgres.Postgres) SegmentRepository {
	return &segmentRepository{
		postgres,
	}
}

func (s *segmentRepository) Create(ctx context.Context, segment string) error {
	query := `INSERT INTO segment (segment_type) VALUES ($1)`

	_, err := s.Pool.Exec(ctx, query, segment)
	return err
}

func (s *segmentRepository) DeleteByID(ctx context.Context, segment *entity.Segment) error {
	query := `DELETE FROM segment WHERE segment_type = $1`

	_, err := s.Pool.Exec(ctx, query, segment.Segment)
	return err
}

func (s *segmentRepository) GetALL(ctx context.Context) ([]entity.Segment, error) {
	query := `SELECT id,segment_type FROM segment`

	rows, err := s.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allSegment := make([]entity.Segment, 0)
	for rows.Next() {
		var segment entity.Segment

		err = rows.Scan(&segment.ID, &segment.Segment)
		if err != nil {
			return nil, err
		}

		allSegment = append(allSegment, segment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return allSegment, nil
}

func (s *segmentRepository) GetByName(ctx context.Context, segmentType string) (*entity.Segment, error) {
	query := `SELECT id,segment_type FROM segment WHERE segment_type = $1`
	segment := &entity.Segment{}

	err := s.Pool.QueryRow(ctx, query, segmentType).Scan(&segment.ID, &segment.Segment)
	if err != nil {
		return nil, err
	}

	return segment, err
}

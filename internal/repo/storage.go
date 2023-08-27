package repo

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
)

type SegmentRepository interface {
	Create(ctx context.Context, segment string) error
	DeleteByID(ctx context.Context, segment *entity.Segment) error
	GetALL(ctx context.Context) ([]entity.Segment, error)
	GetByName(ctx context.Context, segmentType string) (*entity.Segment, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	DeleteByID(ctx context.Context, id string) error
	GetALL(ctx context.Context) (map[string][]string, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	SetSegment(ctx context.Context, user *entity.User) error
	DeleteSegment(ctx context.Context, user *entity.User) error
}

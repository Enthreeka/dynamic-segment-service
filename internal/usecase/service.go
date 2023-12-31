package usecase

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	DeleteUserByID(ctx context.Context, id string) error
	GetAllUser(ctx context.Context) (map[string][]string, error)
	GetUserInfo(ctx context.Context, id string) (*entity.User, error)
	SetSegment(ctx context.Context, user *entity.User) error
	DeleteUserSegment(ctx context.Context, user *entity.User) error
	//GetCSVFile(ctx context.Context, userID string, operation string, date time.Time) (string, error)
}

type SegmentService interface {
	CreateSegment(ctx context.Context, segment string) error
	DeleteSegmentByName(ctx context.Context, segment *entity.Segment) error
	GetIDByName(ctx context.Context, segmentType string) (*entity.Segment, error)
	GetAllSegments(ctx context.Context) ([]entity.Segment, error)
}

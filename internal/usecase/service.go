package usecase

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
)

type UserService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	DeleteUserByID(ctx context.Context, id string) error
	GetAllUser(ctx context.Context) (map[string][]string, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
}

type SegmentService interface {
}

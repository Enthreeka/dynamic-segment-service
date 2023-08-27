package usecase

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/internal/repo"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
)

type userService struct {
	userRepo repo.UserRepository
	log      *logger.Logger
}

func NewUserService(userRepo repo.UserRepository, log *logger.Logger) UserService {
	return &userService{
		userRepo: userRepo,
		log:      log,
	}
}

func (u userService) CreateUser(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) DeleteUserByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetAllUser(ctx context.Context) (map[string][]string, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

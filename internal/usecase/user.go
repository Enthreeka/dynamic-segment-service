package usecase

import (
	"context"
	"errors"
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

func (u *userService) CreateUser(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) DeleteUserByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) GetAllUser(ctx context.Context) (map[string][]string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) GetUserInfo(ctx context.Context, id string) (*entity.User, error) {
	userInfo, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userInfo.Segments == nil {
		return nil, errors.New("zero segments")
	}

	return userInfo, nil
}

func (u *userService) SetSegment(ctx context.Context, user *entity.User) error {
	err := u.userRepo.SetSegment(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) DeleteUserSegment(ctx context.Context, user *entity.User) error {
	err := u.userRepo.DeleteSegment(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

package usecase

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/apperror"
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
	err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		if err == apperror.ErrUserNotFound {
			return nil, apperror.ErrUserNotFound
		}
		apperror.NewAppError(err, "failed to get info about user")
	}

	userInfo, err := u.userRepo.GetSegmentsByUserID(ctx, id)
	if err != nil {
		return nil, apperror.NewAppError(err, "failed to get info about user")
	}

	if userInfo.Segments == nil {
		return nil, apperror.ErrSegmentsNotFound
	}

	return userInfo, nil
}

func (u *userService) SetSegment(ctx context.Context, user *entity.User) error {
	existSegment, err := u.userRepo.GetSegmentsByUserID(ctx, user.ID)
	if len(existSegment.Segments) > 0 {
		u.log.Error("user already exist segments: %v", existSegment.Segments)
		return apperror.ErrUserHasSegment
	}

	err = u.userRepo.SetSegment(ctx, user)
	if err != nil {
		return apperror.NewAppError(err, "failed to set segments to user")
	}

	return nil
}

func (u *userService) DeleteUserSegment(ctx context.Context, user *entity.User) error {
	if len(user.Segments) == 0 {
		return apperror.ErrSegmentsNotFound
	}

	err := u.userRepo.DeleteSegment(ctx, user)
	if err != nil {
		return apperror.NewAppError(err, "failed to delete segments from user")
	}

	return nil
}

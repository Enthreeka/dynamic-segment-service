package usecase

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/apperror"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/internal/repo"
	"github.com/Enthreeka/dynamic-segment-service/pkg/csv"
	"github.com/Enthreeka/dynamic-segment-service/pkg/logger"
	"time"
)

type userService struct {
	userRepo repo.UserRepository
	log      *logger.Logger

	record *csv.Record
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
	users, err := u.userRepo.GetALL(ctx)
	if err != nil {
		if err == apperror.ErrUsersNotFound {
			return nil, apperror.ErrUsersNotFound
		}
		u.log.Error("%v", err)
		return nil, apperror.NewAppError(err, "failed to get all users")
	}

	return users, nil
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
	for _, existEl := range existSegment.Segments {
		for _, newEl := range user.Segments {
			if existEl.Segment == newEl.Segment {
				u.log.Error("user already exist segments: %v", existSegment.Segments)
				return apperror.ErrUserHasSegment
			}
		}
	}

	err = u.userRepo.SetSegment(ctx, user)
	if err != nil {
		return apperror.NewAppError(err, "failed to set segments to user")
	}

	for _, el := range user.Segments {
		r := csv.Record{
			UserID:    user.ID,
			Segment:   el.Segment,
			Operation: "add",
			Date:      time.Now().Format("2006-01-02 15:04:05"),
		}

		r.Write()
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

	for _, el := range user.Segments {
		r := csv.Record{
			UserID:    user.ID,
			Segment:   el.Segment,
			Operation: "delete",
			Date:      time.Now().Format("2006-01-02 15:04:05"),
		}

		r.Write()
	}

	return nil
}

//func (u *userService) GetCSVFile(ctx context.Context, userID string, operation string, date time.Time) (string, error) {
//
//	u.record.Read(userID, operation, date)
//
//}

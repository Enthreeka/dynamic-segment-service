package repo

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/apperror"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	*postgres.Postgres
}

func NewUserReposotory(postgres *postgres.Postgres) UserRepository {
	return &userRepository{
		postgres,
	}
}

func (u *userRepository) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO "user" (id) VALUES (id)`

	_, err := u.Pool.Exec(ctx, query, &user.ID)
	return err
}

func (u *userRepository) DeleteByID(ctx context.Context, id string) error {
	query := `DELETE FROM "user" WHERE id = $1`

	_, err := u.Pool.Exec(ctx, query, id)
	return err
}

func (u *userRepository) GetALL(ctx context.Context) (map[string][]string, error) {
	query := `SELECT "user".id, coalesce(segment.segment_type,'null')
		FROM "user"
        LEFT JOIN user_segment ON "user".id = user_segment.user_id
        LEFT JOIN segment ON segment.id = user_segment.segment_id`

	rows, err := u.Pool.Query(ctx, query)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, apperror.ErrUsersNotFound
		}
		return nil, err
	}

	defer rows.Close()

	usersMap := make(map[string][]string, 0)
	for rows.Next() {
		var id string
		var segment string

		err = rows.Scan(&id, &segment)
		if err != nil {
			return nil, err
		}

		usersMap[id] = append(usersMap[id], segment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usersMap, nil
}

func (u *userRepository) GetSegmentsByUserID(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT segment.segment_type
			FROM "user"
         	JOIN user_segment ON "user".id = user_segment.user_id
         	JOIN segment ON segment.id = user_segment.segment_id
			WHERE "user".id = $1`

	rows, err := u.Pool.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	user := &entity.User{}

	for rows.Next() {
		var segment entity.Segment

		err = rows.Scan(&segment.Segment)
		if err != nil {
			return nil, err
		}

		user.Segments = append(user.Segments, segment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) SetSegment(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO user_segment (user_id,segment_id) VALUES ($1,$2)`

	for _, segment := range user.Segments {
		_, err := u.Pool.Exec(ctx, query, &user.ID, &segment.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *userRepository) DeleteSegment(ctx context.Context, user *entity.User) error {
	query := `DELETE FROM user_segment WHERE user_id = $1 AND segment_id = $2`

	for _, segment := range user.Segments {
		_, err := u.Pool.Exec(ctx, query, user.ID, segment.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *userRepository) FindByID(ctx context.Context, id string) error {
	query := `SELECT id FROM "user" WHERE id = $1`

	err := u.Pool.QueryRow(ctx, query, id).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return apperror.ErrUserNotFound
		}
		return err
	}
	return nil
}

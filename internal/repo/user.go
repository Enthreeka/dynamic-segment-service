package repo

import (
	"context"
	"github.com/Enthreeka/dynamic-segment-service/internal/entity"
	"github.com/Enthreeka/dynamic-segment-service/pkg/postgres"
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
	query := `SELECT "user".id, segment.segment_type
				FROM "user"
         		JOIN user_segment ON "user".id = user_segment.user_id
         		JOIN segment ON segment.id = user_segment.segment_id`

	rows, err := u.Pool.Query(ctx, query)
	if err != nil {
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

func (u *userRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT "user".id, ARRAY_AGG(segment.segment_type) AS segment_types
			FROM "user"
         	JOIN user_segment ON "user".id = user_segment.user_id
         	JOIN segment ON segment.id = user_segment.segment_id
			WHERE "user".id = $1
			GROUP BY "user".id;`
	user := &entity.User{}

	err := u.Pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Segments)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) SetSegment(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO user_segment (user_id,segment_id) VALUES ($1,$2)`

	for _, segment := range user.Segments {
		_, err := u.Pool.Exec(ctx, query, &user.ID, &segment.Segment)
		if err != nil {
			return err
		}
	}
	return nil
}

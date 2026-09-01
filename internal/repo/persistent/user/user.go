package user

import (
	"context"
	"errors"
	"fmt"
	"tenant-gate/internal/entity"
	"tenant-gate/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg}
}

func (r *Repo) Store(ctx context.Context) (string, error) {
	// Implement the logic to store a user in the database using r.pg
	return "", nil
}

func (r *Repo) GetUserByID(ctx context.Context, userID int64) (string, error) {
	// Implement the logic to retrieve a user by ID from the database using r.pg
	return "", nil
}

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	// Implement the logic to retrieve a user by email from the database using r.pg
	return r.getUser(ctx, "email", email)
}

func (r *Repo) getUser(ctx context.Context, column, value string) (entity.User, error) {
	where := squirrel.Eq{}
	if column == "id" {
		uid, err := uuid.Parse(value)
		if err != nil {
			return entity.User{}, fmt.Errorf("UserRepo - getUser - invalid id: %w", err)
		}
		where[column] = uid
	} else {
		where[column] = value
	}

	sql, args, err := r.Builder.
		Select("id, username, email, phone, password_hash, created_at, updated_at").
		From("users").
		Where(where).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - getUser - r.Builder: %w", err)
	}

	var user entity.User

	err = r.Pool.QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, entity.ErrUserNotFound
		}

		return entity.User{}, fmt.Errorf("UserRepo - getUser - r.Pool.QueryRow: %w", err)
	}

	return user, nil
}

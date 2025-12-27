package repository

import (
	"context"
	"database/sql"

	"github.com/TimX-21/auth-service-go/internal/apperror"
	"github.com/TimX-21/auth-service-go/internal/auth/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) getExecutor(c context.Context) DbInterface {
	if c != nil {
		if tx, ok := GetTx(c); ok {
			return tx
		}
	}
	return r.db
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, user model.User) (*model.User, error) {
	conn := r.getExecutor(ctx)

	query := "SELECT id, email, password, is_verified, created_at, updated_at, deleted_at FROM users WHERE email = $1"

	err := conn.QueryRowContext(ctx, query, user.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, apperror.ErrDatabase
	}

	return &user, nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, user model.User) error {
	conn := r.getExecutor(ctx)

	query := "INSERT INTO users (username, email, password, is_verified, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())"

	_, err := conn.ExecContext(ctx, query,
		user.Username,
		user.Email,
		user.Password,
		user.IsActive,
	)
	if err != nil {
		return apperror.ErrDatabase
	}

	return nil
}

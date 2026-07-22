package repository

import (
	"context"

	"task_manager/internal/domain"
	"task_manager/internal/ports"
	"task_manager/internal/repository/db"
)

type userRepository struct {
	queries db.Querier
}

func NewUserRepository(queries db.Querier) ports.UserRepository {
	return &userRepository{
		queries: queries,
	}
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	res, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        u.Email,
		PasswordHash: u.Password,
	})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        res.ID,
		Email:     res.Email,
		Password:  res.PasswordHash,
		CreatedAt: res.CreatedAt.Time,
	}, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	res, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        res.ID,
		Email:     res.Email,
		Password:  res.PasswordHash,
		CreatedAt: res.CreatedAt.Time,
	}, nil
}

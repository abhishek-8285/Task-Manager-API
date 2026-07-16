package ports

import (
	"context"
	"task_manager/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error) 
}
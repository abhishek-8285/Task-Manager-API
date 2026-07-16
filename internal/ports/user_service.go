package ports

import (
	"context"
	"task_manager/internal/domain"
)

type UserService interface {
	Register(ctx context.Context, email,password string) (*domain.User, error)
}
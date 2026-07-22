package service_test

import (
	"context"
	"testing"

	"task_manager/internal/domain"
	"task_manager/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestUserService_Register_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()
	email := "test@example.com"
	password := "securepassword123"

	mockRepo.On("Create", ctx, mock.MatchedBy(func(u *domain.User) bool {
		return u.Email == email && u.Password != password // Hashing verification
	})).Return(&domain.User{
		ID:    1,
		Email: email,
	}, nil)

	createdUser, err := userService.Register(ctx, email, password)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, int64(1), createdUser.ID)
	assert.Equal(t, email, createdUser.Email)

	mockRepo.AssertExpectations(t)
}

func TestUserService_Register_EmptyInput(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()

	createdUser, err := userService.Register(ctx, "", "")

	assert.Error(t, err)
	assert.Nil(t, createdUser)
	assert.Equal(t, "email and password are required", err.Error())
}

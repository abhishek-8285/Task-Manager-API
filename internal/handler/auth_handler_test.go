package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"task_manager/internal/domain"
	"task_manager/internal/handler"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(ctx context.Context, email, password string) (*domain.User, error) {
	args := m.Called(ctx, email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func TestAuthHandler_Register_Success(t *testing.T) {
	mockService := new(MockUserService)
	authHandler := handler.NewAuthHandler(mockService)
	router := handler.SetupRouter(authHandler)

	mockService.On("Register", mock.Anything, "user@example.com", "password123").
		Return(&domain.User{
			ID:    10,
			Email: "user@example.com",
		}, nil)

	body, _ := json.Marshal(handler.RegisterRequest{
		Email:    "user@example.com",
		Password: "password123",
	})

	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var res handler.UserResponse
	err := json.Unmarshal(resp.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), res.ID)
	assert.Equal(t, "user@example.com", res.Email)
}

func TestAuthHandler_Register_ValidationError(t *testing.T) {
	mockService := new(MockUserService)
	authHandler := handler.NewAuthHandler(mockService)
	router := handler.SetupRouter(authHandler)

	// Invalid email and password length
	body, _ := json.Marshal(map[string]string{
		"email":    "invalid-email",
		"password": "123",
	})

	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

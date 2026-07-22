package handler

import (
	"net/http"

	"task_manager/internal/ports"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService ports.UserService
}

func NewAuthHandler(userService ports.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:    user.ID,
		Email: user.Email,
	})
}
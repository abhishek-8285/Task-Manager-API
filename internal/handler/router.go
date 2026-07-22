package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *AuthHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
		}
	}

	return r
}

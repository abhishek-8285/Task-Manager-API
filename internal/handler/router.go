package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	healthHandler := NewHealthHandler()

	router.GET("/health", healthHandler.Health)

	return router
}

package main

import (
	"log"

	"task_manager/internal/handler"
	"task_manager/internal/repository"
	"task_manager/internal/service"
)

func main() {
	// 1. Instantiation of layers (Hexagonal Dependency Injection)
	userRepo := repository.NewUserRepository(nil) // DB pool will be attached in full docker setup
	userService := service.NewUserService(userRepo)
	authHandler := handler.NewAuthHandler(userService)

	// 2. Setup HTTP Router with AuthHandler
	router := handler.SetupRouter(authHandler)

	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

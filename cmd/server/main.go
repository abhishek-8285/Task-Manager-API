package main

import (
	"log"
	"task_manager/internal/handler"
)

func main() {
	router := handler.SetupRouter()
	log.Println("Server started on :8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}

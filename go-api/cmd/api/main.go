package main

import (
	"log"
	"net/http"

	"go-api/internal/handler"
	"go-api/internal/repository"
	"go-api/internal/service"
)

func main() {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("GET /users/{id}", userHandler.FindByID)

	log.Println("http://localhost:8080")

	log.Fatal(
		http.ListenAndServe(":8080", mux),
	)
}

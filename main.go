package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"log"
	"net/http"
	"user-api/domain"
	"user-api/repository"
	"user-api/server"
)

func main() {
	s := server.Server{
		Repository: &repository.UserInMemoryRepository{
			Users: make(map[uuid.UUID]domain.User),
		},
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/users", s.CreateUser)
	r.Get("/users/{id}", s.GetUser)
	r.Get("/users", s.GetUsers)
	r.Patch("/users/{id}", s.UpdateUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}

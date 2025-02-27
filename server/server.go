package server

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"user-api/domain"
	"user-api/repository"
)

type Server struct {
	Repository repository.UserRepository
}

func (s Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	cur := CreateUserRequest{}
	if err := json.Unmarshal(body, &cur); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if cur.Login == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := domain.User{
		Id:    uuid.New(),
		Login: cur.Login,
	}

	userId, err := s.Repository.CreateUser(u)
	if err != nil {
		log.Println("Failed to create user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := CreateUserResponse{
		ID: userId,
	}

	payload, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed to marshal response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}
func (s Server) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Failed to parse userId", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.Repository.GetUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to get user:", err)
		return
	}

	isNotFound := user == domain.User{}
	if isNotFound {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("User %d not found", userId)
		return
	}

	gur := GetUserResponse{Login: user.Login}
	payload, err := json.Marshal(gur)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to marshal response:", err)
	}

	w.Write(payload)
}
func (s Server) GetUsers(w http.ResponseWriter, _ *http.Request) {

	users, err := s.Repository.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to get user:", err)
		return
	}

	var userLogins []string
	for _, user := range users {
		userLogins = append(userLogins, user.Login)
	}

	gur := GetUsersResponse{Logins: userLogins}
	payload, err := json.Marshal(gur)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Failed to marshal response:", err)
	}

	w.Write(payload)
}
func (s Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Failed to parse userId", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	cur := UpdateUserRequest{}
	if err := json.Unmarshal(body, &cur); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if cur.Login == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u := domain.User{
		Id:    userId,
		Login: cur.Login,
	}

	updatedUserId, err := s.Repository.UpdateUser(u)
	if err != nil {
		log.Println("Failed to create user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := UpdateUserResponse{
		ID: updatedUserId,
	}

	payload, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed to marshal response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(payload)
}

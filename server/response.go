package server

import "github.com/google/uuid"

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type UpdateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetUserResponse struct {
	Login string `json:"login"`
}
type GetUsersResponse struct {
	Logins []string `json:"logins"`
}

package server

type CreateUserRequest struct {
	Login string `json:"login"`
}

type UpdateUserRequest struct {
	Login string `json:"login"`
}

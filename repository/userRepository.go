package repository

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"user-api/domain"
)

type UserRepository interface {
	CreateUser(u domain.User) (uuid.UUID, error)
	GetUser(userId uuid.UUID) (domain.User, error)
	GetUsers() ([]domain.User, error)
	UpdateUser(u domain.User) (uuid.UUID, error)
}

type UserInMemoryRepository struct {
	Users    map[uuid.UUID]domain.User
	userLock sync.RWMutex
}

func (r *UserInMemoryRepository) CreateUser(u domain.User) (uuid.UUID, error) {
	r.userLock.Lock()
	defer r.userLock.Unlock()

	r.Users[u.Id] = u
	return u.Id, nil
}

func (r *UserInMemoryRepository) GetUser(userId uuid.UUID) (domain.User, error) {
	r.userLock.RLock()
	defer r.userLock.RUnlock()

	u, _ := r.Users[userId]
	return u, nil
}

func (r *UserInMemoryRepository) UpdateUser(u domain.User) (uuid.UUID, error) {
	r.userLock.Lock()
	defer r.userLock.Unlock()

	_, ok := r.Users[u.Id]
	if !ok {
		return uuid.UUID{}, errors.New("user not found")
	}
	r.Users[u.Id] = u
	return u.Id, nil
}

func (r *UserInMemoryRepository) GetUsers() ([]domain.User, error) {
	r.userLock.RLock()
	defer r.userLock.RUnlock()

	var users []domain.User
	for _, user := range r.Users {
		users = append(users, user)
	}

	return users, nil
}

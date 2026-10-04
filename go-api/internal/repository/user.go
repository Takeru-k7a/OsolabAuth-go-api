package repository

import (
	"errors"

	"go-api/internal/model"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository struct {
	users  map[int]model.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

func (r *UserRepository) Create(user model.User) (model.User, error) {
	user.ID = r.nextID
	r.nextID++

	r.users[user.ID] = user

	return user, nil
}

func (r *UserRepository) FindByID(id int) (model.User, error) {
	user, exists := r.users[id]

	if !exists {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
}

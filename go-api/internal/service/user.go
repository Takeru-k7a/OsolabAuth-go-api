package service

import (
	"errors"
	"strings"

	"go-api/internal/model"
	"go-api/internal/repository"
)

var ErrNameRequired = errors.New("name is required")

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(
	repository *repository.UserRepository,
) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Create(
	name string,
	email string,
) (model.User, error) {

	name = strings.TrimSpace(name)

	if name == "" {
		return model.User{}, ErrNameRequired
	}

	user := model.User{
		Name:  name,
		Email: email,
	}

	return s.repository.Create(user)
}

func (s *UserService) FindByID(id int) (model.User, error) {
	return s.repository.FindByID(id)
}

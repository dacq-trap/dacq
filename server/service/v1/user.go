package v1

import (
	"github.com/dacq-trap/dacq/server/model"
	"github.com/dacq-trap/dacq/server/repository"
	"github.com/dacq-trap/dacq/server/service"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) service.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) ReadUserByName(name string) (*model.User, error) {
	return nil, nil
}

func (s *userService) CreateUser(name string) (*model.User, error) {
	return nil, nil
}

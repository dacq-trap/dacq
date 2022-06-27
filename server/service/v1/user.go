package v1

import (
	"github.com/dacq-trap/dacq/server/model"
	"github.com/dacq-trap/dacq/server/repository"
	"github.com/dacq-trap/dacq/server/service"
)

type usersService struct {
	repo repository.UsersRepository
}

func NewUsersService(repo repository.UsersRepository) service.UsersService {
	return &usersService{
		repo: repo,
	}
}

func (s *usersService) ReadUserByName(name string) (*model.User, error) {
	return nil, nil
}

func (s *usersService) CreateUser(name string) (*model.User, error) {
	return nil, nil
}

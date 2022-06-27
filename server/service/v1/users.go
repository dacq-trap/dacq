package v1

import (
	"context"

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

func (s *usersService) ReadUserByName(ctx context.Context, name string) (*model.User, error) {
	return s.repo.SelectUserByName(ctx, name)
}

func (s *usersService) CreateUser(ctx context.Context, name string) (*model.User, error) {
	url := "https://q.trap.jp/api/v3/public/icon/" + name

	user := model.User{
		Name:    name,
		IconURL: url,
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

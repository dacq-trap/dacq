package users

import "github.com/dacq-trap/dacq/server/service"

type UsersHandler struct {
	service service.UsersService
}

func NewUsersHandler(service service.UsersService) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

package users

import "github.com/dacq-trap/dacq/server/service"

type UsersHandler struct {
	service service.UserService
}

func NewUsersHandler(service service.UserService) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

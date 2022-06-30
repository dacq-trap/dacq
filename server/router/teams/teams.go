package teams

import "github.com/dacq-trap/dacq/server/service"

type TeamsHandler struct {
	service service.TeamsService
}

func NewTeamsHandler(service service.TeamsService) *TeamsHandler {
	return &TeamsHandler{
		service: service,
	}
}

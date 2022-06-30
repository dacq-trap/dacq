package teams

import (
	"github.com/dacq-trap/dacq/server/service"
	"github.com/google/uuid"
)

type TeamCoreInfo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type TeamsHandler struct {
	service service.TeamsService
}

func NewTeamsHandler(service service.TeamsService) *TeamsHandler {
	return &TeamsHandler{
		service: service,
	}
}

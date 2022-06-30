package teams

import (
	"errors"
	"net/http"

	"github.com/dacq-trap/dacq/server/model"
	"github.com/dacq-trap/dacq/server/router/session"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

/*
type: object
description: チームの基本情報
properties:
  id:
    type: string
    format: uuid
    description: ID
  name:
    type: string
    description: チーム名
required:
  - id
  - name
*/

type TeamCoreInfo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetMyTeamsResponse struct {
	TeamCoreInfos []TeamCoreInfo `json:"teams"`
}

func (h *TeamsHandler) GetMyTeams(c echo.Context) error {
	username := c.Get(session.UserNameKey).(string)

	teams, err := h.service.ReadTeamsByUserName(c.Request().Context(), username)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := GetMyTeamsResponse{
		TeamCoreInfos: make([]TeamCoreInfo, len(teams)),
	}
	for i, team := range teams {
		res.TeamCoreInfos[i] = TeamCoreInfo{
			ID:   team.ID,
			Name: team.Name,
		}
	}
	return c.JSON(http.StatusOK, res)
}

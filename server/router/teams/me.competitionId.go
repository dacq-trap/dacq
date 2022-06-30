package teams

import (
	"errors"
	"net/http"

	"github.com/dacq-trap/dacq/server/model"
	"github.com/dacq-trap/dacq/server/router/session"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *TeamsHandler) GetMyTeamInCompetition(c echo.Context) error {
	competitionId, err := uuid.Parse(c.Param("competitionId"))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	username := c.Get(session.UserNameKey).(string)

	team, err := h.service.ReadTeamByUserNameAndCompetitionID(c.Request().Context(), username, competitionId)
	if err != nil {
		if errors.Is(err, model.ErrUnauthorized) {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := TeamCoreInfo{
		ID:   team.ID,
		Name: team.Name,
	}
	return c.JSON(http.StatusOK, res)
}

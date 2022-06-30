package teams

import (
	"errors"
	"net/http"
	"time"

	"github.com/dacq-trap/dacq/server/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// schema/openapi/competitonCoreInfo.yaml
type CompetitionCoreInfo struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}

// schema/openapi/common/user.yaml
type User struct {
	Name    string `json:"name"`
	IconUrl string `json:"iconUrl"`
}

// schema/openapi/common/team.yaml
type GetTeamResponse struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Competition CompetitionCoreInfo `json:"competition"`
	Users       []User              `json:"users"`
	IsMerging   bool                `json:"isMerging"`
}

func (h *TeamsHandler) GetTeam(c echo.Context) error {
	teamId, err := uuid.Parse(c.Param("teamId"))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	team, err := h.service.ReadTeamByID(c.Request().Context(), teamId)
	if err != nil {
		if errors.Is(err, model.ErrUnauthorized) {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := GetTeamResponse{
		ID:   team.ID,
		Name: team.Name,
		Competition: CompetitionCoreInfo{
			ID:      team.Competition.ID,
			Name:    team.Competition.Name,
			StartAt: team.Competition.StartAt,
			EndAt:   team.Competition.EndAt,
		},
		Users:     make([]User, len(team.Users)),
		IsMerging: team.IsMerging,
	}
	for i, user := range team.Users {
		res.Users[i] = User{
			Name:    user.Name,
			IconUrl: user.IconURL,
		}
	}
	return c.JSON(http.StatusOK, res)
}

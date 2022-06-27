package users

import (
	"errors"
	"net/http"

	"github.com/dacq-trap/dacq/server/model"
	"github.com/dacq-trap/dacq/server/router/session"
	"github.com/labstack/echo/v4"
)

type GetMeResponse struct {
	Name    string `json:"name"`
	IconUrl string `json:"iconUrl"`
}

func (h *UsersHandler) GetMe(c echo.Context) error {
	name := c.Request().Context().Value(session.UserNameKey).(string)

	user, err := h.service.ReadUserByName(c.Request().Context(), name)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := GetMeResponse{
		Name:    user.Name,
		IconUrl: user.IconURL,
	}
	return c.JSON(http.StatusOK, res)
}

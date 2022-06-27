package users

import (
	"errors"
	"net/http"

	"github.com/dacq-trap/dacq/server/model"
	"github.com/dacq-trap/dacq/server/router/session"
	"github.com/labstack/echo/v4"
)

func (h *UsersHandler) GetMe(c echo.Context) error {
	name := c.Request().Context().Value(session.UserNameKey).(string)

	user, err := h.service.ReadUserByName(name)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}

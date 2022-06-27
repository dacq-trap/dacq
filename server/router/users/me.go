package users

import (
	"github.com/dacq-trap/dacq/server/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *UsersHandler) GetMe(c echo.Context) error {
	name := c.Request().Context().Value(session.UserNameKey).(string)
	user, err := h.service.ReadUserByName(name)
	if err == model.ErrNotFound {
		return echo.NewHTTPError(http.StatusNotFound)
	} else if err == model.ErrForbidden {
		return echo.NewHTTPError(http.StatusForbidden)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, user)
}

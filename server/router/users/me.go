package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

/*
type: object
properties:
  name:
    type: string
    description: ユーザー名 (traP ID)
    example: "trapyojo"
  iconUrl:
    type: string
    description: アイコンURL
    example: https://q.trap.jp/api/v3/public/icon/trapyojo
required:
  - name
  - iconUrl
*/

type user struct {
	Name    string `json:"name"`
	IconUrl string `json:"iconUrl"`
}

func (h *UsersHandler) GetMe(c echo.Context) error {
	name := c.Param("UserName")
	user := user{
		Name:    name,
		IconUrl: "https://q.trap.jp/api/v3/public/icon/" + name,
	}
	return c.JSON(http.StatusOK, user)
}

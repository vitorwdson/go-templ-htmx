package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/vitorwdson/go-templ-htmx/utils"
)

func (h Handler) Logout(c echo.Context) error {
	session, err := GetSession(c, h.Redis, h.DB)
	if err == nil {
		KillSession(c, h.Redis, *session)
	}

	return utils.RedirectHtmx(c, "/login")
}

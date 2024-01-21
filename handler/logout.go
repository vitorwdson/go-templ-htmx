package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/vitorwdson/go-templ-htmx/utils"
)

func (h Handler) Logout(c echo.Context) error {
	session, err := h.GetSession(c)
	if err == nil {
		h.KillSession(c, *session)
	}

	return utils.RedirectHtmx(c, "/login")
}

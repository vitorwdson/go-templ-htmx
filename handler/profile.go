package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/vitorwdson/go-templ-htmx/utils"
	userView "github.com/vitorwdson/go-templ-htmx/view/user"
)

func (h Handler) Profile(c echo.Context) error {
	session, err := h.GetSession(c)
	if err != nil {
		return utils.RedirectHtmx(c, "/login")
	}

	return utils.Render(c, userView.Profile(session.User))
}

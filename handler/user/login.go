package user

import (
	"github.com/labstack/echo/v4"
	"github.com/vitorwdson/go-templ-htmx/utils"
	userView "github.com/vitorwdson/go-templ-htmx/view/user"
)

func (h UserHandler) Login(c echo.Context) error {
	return utils.Render(c, userView.Login())
}


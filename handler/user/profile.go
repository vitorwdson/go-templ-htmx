package user

import (
	"github.com/labstack/echo/v4"
	userModel "github.com/vitorwdson/go-templ-htmx/model/user"
	"github.com/vitorwdson/go-templ-htmx/utils"
	userView "github.com/vitorwdson/go-templ-htmx/view/user"
)

func (h UserHandler) Profile(c echo.Context) error {
	u := userModel.User{
		Name: "Vitor Wdson",
		Email: "vitor.wdson2@gmail.com",
	}

	return utils.Render(c, userView.Profile(u))
}



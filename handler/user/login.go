package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	userModel "github.com/vitorwdson/go-templ-htmx/model/user"
	"github.com/vitorwdson/go-templ-htmx/utils"
	userView "github.com/vitorwdson/go-templ-htmx/view/user"
)

type loginFormData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (h UserHandler) Login(c echo.Context) error {
	props := userView.LoginViewProps{}
	return utils.Render(c, userView.Login(props))
}

func (h UserHandler) PostLogin(c echo.Context) error {
	var data registerFormData

	err := c.Bind(&data)
	if err != nil {
		return err
	}

	user, err := userModel.GetByUsername(h.DB, data.Username)
	if err != nil || user == nil || !user.ValidatePassword(data.Password) {
		props := userView.LoginViewProps{
			Username: data.Username,
			Error:    "The informed username and/or password is incorrect",
		}

		return utils.Render(c, userView.Login(props))
	}

	err = h.authenticateUser(c, user)
	if err != nil {
		return err
	}

	return utils.RedirectHtmx(c, "/profile")

}

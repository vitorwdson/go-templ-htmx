package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	userModel "github.com/vitorwdson/go-templ-htmx/model/user"
	"github.com/vitorwdson/go-templ-htmx/utils"
	userView "github.com/vitorwdson/go-templ-htmx/view/user"
)

type registerFormData struct {
	Name            string `form:"name"`
	Username        string `form:"username"`
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm-password"`
}

func (h UserHandler) Register(c echo.Context) error {
	props := userView.RegisterViewProps{}

	return utils.Render(c, userView.Register(props))
}

func (h UserHandler) PostRegister(c echo.Context) error {
	var data registerFormData

	err := c.Bind(&data)
	if err != nil {
		return err
	}

	nameError := ""
	if data.Name == "" {
		nameError = "The user does not have a name"
	}

	usernameError := ""
	if data.Username == "" {
		usernameError = "The user does not have an username"
	} else if len(data.Username) > 30 {
		usernameError = "The username is too big (max: 30)"
	} else if u, _ := userModel.GetByUsername(h.DB, data.Username); u != nil {
		usernameError = "This username already exists."
	}

	passwordError := userModel.CheckPasswordStrength(data.Password)
	if passwordError == "" && data.Password != data.ConfirmPassword {
		passwordError = "Both passwords must match."
	}

	emailError := ""
	if data.Email == "" {
		emailError = "The user does not have a email"
	} else if len(data.Email) > 100 {
		emailError = "The nickname is too big (max: 100)"
	}

	props := userView.RegisterViewProps{
		Name:          data.Name,
		Username:      data.Username,
		Email:         data.Email,
		NameError:     nameError,
		UsernameError: usernameError,
		EmailError:    emailError,
		PasswordError: passwordError,
	}

	return utils.Render(c, userView.Register(props))
}

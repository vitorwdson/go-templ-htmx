package handler

import (
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

func (h Handler) Register(c echo.Context) error {
	props := userView.RegisterViewProps{}

	return utils.Render(c, userView.Register(props))
}

func (h Handler) PostRegister(c echo.Context) error {
	var data registerFormData

	err := c.Bind(&data)
	if err != nil {
		return err
	}

	validation := userModel.Validate(
		data.Name,
		data.Username,
		data.Email,
		data.Password,
		data.ConfirmPassword,
		h.DB,
	)

	var user *userModel.User
	if !validation.Error {
		user = &userModel.User{
			Name:     data.Name,
			Username: data.Username,
			Email:    data.Email,
		}

		err := user.SetPassword([]byte(data.Password))
		if err != nil {
			return err
		}

		err = user.Save(h.DB)
		if err != nil {
			return err
		}

		err = h.authenticateUser(c, *user)
		if err != nil {
			return err
		}

		return utils.RedirectHtmx(c, "/profile")
	}

	props := userView.RegisterViewProps{
		Name:          data.Name,
		Username:      data.Username,
		Email:         data.Email,
		NameError:     validation.NameError,
		UsernameError: validation.UsernameError,
		EmailError:    validation.EmailError,
		PasswordError: validation.PasswordError,
	}

	return utils.Render(c, userView.Register(props))
}

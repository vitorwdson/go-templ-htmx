package user

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

	validationError := false

	nameError := ""
	if data.Name == "" {
		nameError = "The user does not have a name"
		validationError = true
	}

	usernameError := ""
	if data.Username == "" {
		usernameError = "The user does not have an username"
		validationError = true
	} else if len(data.Username) > 30 {
		usernameError = "The username is too big (max: 30)"
		validationError = true
	} else if u, _ := userModel.GetByUsername(h.DB, data.Username); u != nil {
		usernameError = "This username already exists."
		validationError = true
	}

	passwordError := userModel.CheckPasswordStrength(data.Password)
	if passwordError == "" && data.Password != data.ConfirmPassword {
		passwordError = "Both passwords must match."
	}

	if passwordError != "" {
		validationError = true
	}

	emailError := ""
	if data.Email == "" {
		emailError = "The user does not have a email"
		validationError = true
	} else if len(data.Email) > 100 {
		emailError = "The nickname is too big (max: 100)"
		validationError = true
	}

	var user *userModel.User
	if !validationError {
		user = &userModel.User{
			Name:     data.Name,
			Username: data.Username,
			Email:    data.Email,
		}

		err := user.SetPassword([]byte(data.Password))
		if err != nil {
			passwordError = err.Error()
			user = nil
		}
	}

	if user != nil {
		err := user.Save(h.DB)
		if err != nil {
			return err
		}

		err = h.authenticateUser(c, user)
		if err != nil {
			return err
		}

		return utils.RedirectHtmx(c, "/profile")
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

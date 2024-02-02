package handler

import (
	"net/http"

	"github.com/vitorwdson/go-templ-htmx/data/models"
	"github.com/vitorwdson/go-templ-htmx/utils"
	"github.com/vitorwdson/go-templ-htmx/view/pages"
)

func (s server) handleRegister(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleRegisterGET(w, r)
	} else if r.Method == http.MethodPost {
		return s.handleRegisterPOST(w, r)
	}

	return InvalidMethod
}

func (s server) handleRegisterGET(w http.ResponseWriter, r *http.Request) error {
	props := pages.RegisterViewProps{}

	return utils.Render(w, r, pages.Register(props))
}

func (s server) handleRegisterPOST(w http.ResponseWriter, r *http.Request) error {
	data, err := parseRegisterFormData(r)
	if err != nil {
		return err
	}
	s.Logger.Println(data)

	validation := s.UserRepo.Validate(
		data.Name,
		data.Username,
		data.Email,
		data.Password,
		data.ConfirmPassword,
	)

	var user *models.User
	if !validation.Error {
		user = &models.User{
			Name:     data.Name,
			Username: data.Username,
			Email:    data.Email,
		}

		err := user.SetPassword([]byte(data.Password))
		if err != nil {
			return err
		}

		err = s.UserRepo.Save(user)
		if err != nil {
			return err
		}

		// err = s.authenticateUser(w, r, *user)
		// if err != nil {
		// 	return err
		// }

		return utils.RedirectHtmx(w, r, "/profile")
	}

	props := pages.RegisterViewProps{
		Name:          data.Name,
		Username:      data.Username,
		Email:         data.Email,
		NameError:     validation.NameError,
		UsernameError: validation.UsernameError,
		EmailError:    validation.EmailError,
		PasswordError: validation.PasswordError,
	}

	return utils.Render(w, r, pages.Register(props))
}

type registerFormData struct {
	Name            string
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}

func parseRegisterFormData(r *http.Request) (*registerFormData, error) {
	r.ParseForm()

	data := registerFormData{
		Name:            r.Form.Get("name"),
		Username:        r.Form.Get("username"),
		Email:           r.Form.Get("email"),
		Password:        r.Form.Get("password"),
		ConfirmPassword: r.Form.Get("confirm-password"),
	}

	return &data, nil
}

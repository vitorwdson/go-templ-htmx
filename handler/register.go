package handler

import (
	"database/sql"
	"net/http"

	"github.com/vitorwdson/go-templ-htmx/data/validation"
	"github.com/vitorwdson/go-templ-htmx/db"
	"github.com/vitorwdson/go-templ-htmx/utils"
	"github.com/vitorwdson/go-templ-htmx/view/pages"
)

func (s server) handleRegister(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleRegisterGET(w, r)
	} else if r.Method == http.MethodPost {
		return s.handleRegisterPOST(w, r)
	}

	return InvalidMethodError
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

	result := validation.ValidateUser(
		data.Name,
		data.Username,
		data.Email,
		data.Password,
		data.ConfirmPassword,
		r.Context(),
		s.DB,
	)

	if !result.Error {
		passwordHash, err := validation.SetPassword([]byte(data.Password))
		if err != nil {
			return err
		}

		user, err := s.DB.CreateUser(r.Context(), db.CreateUserParams{
			Name:     data.Name,
			Username: data.Username,
			PasswordHash: passwordHash,
			Email:    sql.NullString{String: data.Email, Valid: true},
		})
		if err != nil {
			return err
		}

		err = s.authenticateUser(w, r, user)
		if err != nil {
			return err
		}

		return utils.RedirectHtmx(w, r, "/profile")
	}

	props := pages.RegisterViewProps{
		Name:          data.Name,
		Username:      data.Username,
		Email:         data.Email,
		NameError:     result.NameError,
		UsernameError: result.UsernameError,
		EmailError:    result.EmailError,
		PasswordError: result.PasswordError,
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

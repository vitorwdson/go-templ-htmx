package handler

import (
	"net/http"

	"github.com/vitorwdson/go-templ-htmx/utils"
	"github.com/vitorwdson/go-templ-htmx/validation"
	"github.com/vitorwdson/go-templ-htmx/view/pages"
)

func (s server) handleLoginGET(w http.ResponseWriter, r *http.Request) error {
	props := pages.LoginViewProps{}

	return utils.Render(w, r, pages.Login(props))
}

func (s server) handleLoginPOST(w http.ResponseWriter, r *http.Request) error {
	data, err := parseLoginFormData(r)
	if err != nil {
		return err
	}

	user, err := s.DB.GetUserByUsername(r.Context(), data.Username)
	if err != nil || !validation.ValidatePassword(user, data.Password) {
		props := pages.LoginViewProps{
			Username: data.Username,
			Error:    "The informed username and/or password is incorrect",
		}

		return utils.Render(w, r, pages.Login(props))
	}

	err = s.authenticateUser(w, r, user)
	if err != nil {
		return err
	}

	return utils.RedirectHtmx(w, r, "/profile")
}

type loginFormData struct {
	Username string
	Password string
}

func parseLoginFormData(r *http.Request) (*loginFormData, error) {
	r.ParseForm()

	data := loginFormData{
		Username: r.Form.Get("username"),
		Password: r.Form.Get("password"),
	}

	return &data, nil
}

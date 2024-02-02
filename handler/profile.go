package handler

import (
	"net/http"

	"github.com/vitorwdson/go-templ-htmx/utils"
	"github.com/vitorwdson/go-templ-htmx/view/pages"
)

func (s server) handleProfile(w http.ResponseWriter, r *http.Request) error {
	session, err := s.GetSession(w, r)
	if err != nil {
		return UserNotAuthenticatedError
	}

	return utils.Render(w, r, pages.Profile(session.User.Name))
}

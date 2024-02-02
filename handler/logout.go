package handler

import (
	"net/http"

	"github.com/vitorwdson/go-templ-htmx/utils"
)

func (s server) handleLogout(w http.ResponseWriter, r *http.Request) error {
	session, err := s.GetSession(w, r)
	if err == nil {
		s.KillSession(w, r, *session)
	}

	return utils.RedirectHtmx(w, r, "/login")
}

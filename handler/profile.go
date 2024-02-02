package handler

import (
	"net/http"

	"github.com/vitorwdson/go-templ-htmx/utils"
	"github.com/vitorwdson/go-templ-htmx/view/pages"
)

func (s server) handleProfile(w http.ResponseWriter, r *http.Request) error {
	// session, err := h.GetSession(c)
	// if err != nil {
	// 	return utils.RedirectHtmx(c, "/login")
	// }

	return utils.Render(w, r, pages.Profile("John Doe"))
}

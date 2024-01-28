package handler

import (
	"net/http"

	"github.com/vitorwdson/go-templ-htmx/utils"
)

type ServerError struct {
	Message    string
	StatusCode int
}

func (e ServerError) Error() string {
	return e.Message
}

var (
	UserNotAuthenticated = ServerError{
		Message:    "User not authenticated",
		StatusCode: http.StatusForbidden,
	}

	UserNotAuthorizedError = ServerError{
		Message:    "User not allowed",
		StatusCode: http.StatusForbidden,
	}

	InvalidBody = ServerError{
		Message:    "Invalid body",
		StatusCode: http.StatusBadRequest,
	}
)

func (h handler) handleErrors(f RouteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}

		serr, ok := err.(ServerError)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		if serr == UserNotAuthenticated {
			utils.RedirectHtmx(w, r, "/login")
			return
		}

		w.WriteHeader(serr.StatusCode)
		w.Write([]byte(serr.Message))
	}
}

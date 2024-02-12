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
	UserNotAuthenticatedError = ServerError{
		Message:    "User not authenticated",
		StatusCode: http.StatusForbidden,
	}

	UserNotAuthorizedError = ServerError{
		Message:    "User not allowed",
		StatusCode: http.StatusForbidden,
	}

	InvalidBodyError = ServerError{
		Message:    "Invalid body",
		StatusCode: http.StatusBadRequest,
	}

	InvalidMethodError = ServerError{
		Message:    "Method not allowed",
		StatusCode: http.StatusMethodNotAllowed,
	}
)

func (s server) handleErrors(f RouteHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Println(r.Method, r.URL.Path)

		err := f(w, r)
		if err == nil {
			return
		}

		s.Logger.Println("Error: ", err)

		serr, ok := err.(ServerError)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		if serr == UserNotAuthenticatedError {
			utils.RedirectHtmx(w, r, "/login")
			return
		}

		w.WriteHeader(serr.StatusCode)
		w.Write([]byte(serr.Message))
	}
}

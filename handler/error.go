package handler

import "net/http"

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

		// TODO: Handle errors here
	}
}

package validation

import (
	"context"
	"strings"

	"github.com/vitorwdson/go-templ-htmx/db"
)

type validationResult struct {
	Error         bool
	NameError     string
	UsernameError string
	EmailError    string
	PasswordError string
}

func ValidateUser(
	name, username, email, password, confirmPassword string,
	ctx context.Context,
	db *db.Queries,
) validationResult {
	result := validationResult{
		Error: false,
	}

	if name == "" {
		result.NameError = "The user does not have a name"
		result.Error = true
	}

	if username == "" {
		result.UsernameError = "The user does not have an username"
		result.Error = true
	} else if len(username) > 30 {
		result.UsernameError = "The username is too big (max: 30)"
		result.Error = true
	} else if _, err := db.GetUserByUsername(ctx, username); err == nil {
		result.UsernameError = "This username already exists."
		result.Error = true
	}

	result.PasswordError = CheckPasswordStrength(password)
	if len([]byte(password)) > 72 {
		result.PasswordError = "The informed password is too big"
	} else if result.PasswordError == "" && password != confirmPassword {
		result.PasswordError = "Both passwords must match."
	}

	if result.PasswordError != "" {
		result.Error = true
	}

	if email == "" {
		result.EmailError = "The user does not have a email"
		result.Error = true
	} else if len(email) > 100 {
		result.EmailError = "The nickname is too big (max: 100)"
		result.Error = true
	}

	return result
}

func CheckPasswordStrength(password string) string {
	if len(password) < 8 {
		return "The password must contain at least eight characters."
	}

	if !strings.ContainsFunc(password, func(r rune) bool {
		return r >= '0' && r <= '9'
	}) {
		return "The password must contain at least one number."
	}

	if !strings.ContainsFunc(password, func(r rune) bool {
		return r >= 'A' && r <= 'Z'
	}) {
		return "The password must contain at least one uppercase letter."
	}

	if !strings.ContainsFunc(password, func(r rune) bool {
		return r >= 'a' && r <= 'z'
	}) {
		return "The password must contain at least one lowercase letter."
	}

	return ""
}

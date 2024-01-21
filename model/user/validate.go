package user

import (
	"database/sql"
)

type validationResult struct {
	Error         bool
	NameError     string
	UsernameError string
	EmailError    string
	PasswordError string
}

func Validate(
	name, username, email, password, confirmPassword string,
	db *sql.DB,
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
	} else if u, _ := GetByUsername(db, username); u != nil {
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

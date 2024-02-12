package validation

import (
	"errors"

	"github.com/vitorwdson/go-templ-htmx/db"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(u db.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func SetPassword(newPassword []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Could not hash password")
	}

	return string(hashedPassword), nil
}

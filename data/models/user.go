package models

import (
	"errors"

	"github.com/vitorwdson/go-templ-htmx/db"
	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(u db.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func SetPassword(u *db.User, newPassword []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Could not hash password")
	}

	u.PasswordHash = string(hashedPassword)
	return nil
}

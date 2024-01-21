package user

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) SetPassword(newPassword []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Could not hash password")
	}

	u.password = string(hashedPassword)
	return nil
}

func (u User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
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

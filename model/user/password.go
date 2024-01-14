package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (u *User) SetPassword(newPassword []byte) error {
	if len(newPassword) > 72 {
		return errors.New("The informed password is too big")
	}

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
	if password == "" {
		return "The password can not be empty."
	}

	return ""
}

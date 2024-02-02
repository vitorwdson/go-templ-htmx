package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Name     string
	Email    string
	Password string
}

func (u User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) SetPassword(newPassword []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(newPassword, bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Could not hash password")
	}

	u.Password = string(hashedPassword)
	return nil
}

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


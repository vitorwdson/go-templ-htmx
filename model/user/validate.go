package user

import "errors"

func (u User) Validate() error {
	if u.Name == "" {
		return errors.New("The user does not have a name")
	}

	if u.Username == "" {
		return errors.New("The user does not have an username")
	}
	if len(u.Username) > 30 {
		return errors.New("The username is too big (max: 30)")
	}

	if u.ID == 0 && u.password == "" {
		return errors.New("The user does not have a password")
	}

	if u.Email == "" {
		return errors.New("The user does not have a email")
	}
	if len(u.Email) > 100 {
		return errors.New("The nickname is too big (max: 100)")
	}

	return nil
}

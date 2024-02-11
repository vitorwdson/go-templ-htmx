package repos

import (
	"context"

	"github.com/vitorwdson/go-templ-htmx/db"
)

type UserRepo struct {
	DB *db.Queries
}

func NewUserRepo(db *db.Queries) UserRepo {
	return UserRepo{
		DB: db,
	}
}

func (r UserRepo) Save(u *db.User) error {
	if u.ID != 0 {
		// User exists in db, should update
		err := r.DB.UpdateUser(context.Background(), db.UpdateUserParams{
			PasswordHash: u.PasswordHash,
			Name:         u.Name,
			Email:        u.Email,
			ID:           u.ID,
		})
		if err != nil {
			return err
		}
	} else {
		// User doesn't exists in db, should insert
		newId, err := r.DB.CreateUser(context.Background(), db.CreateUserParams{
			Username:     u.Username,
			PasswordHash: u.PasswordHash,
			Name:         u.Name,
			Email:        u.Email,
		})
		if err != nil {
			return err
		}

		u.ID = newId
	}

	return nil
}

func (r UserRepo) GetByID(id int32) (*db.User, error) {
	user, err := r.DB.GetUserByID(context.Background(), id)
	return &user, err
}

func (r UserRepo) GetByUsername(username string) (*db.User, error) {
	user, err := r.DB.GetUserByUsername(context.Background(), username)
	return &user, err
}

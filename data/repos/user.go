package repos

import (
	"database/sql"

	"github.com/vitorwdson/go-templ-htmx/data/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return UserRepo{
		DB: db,
	}
}

func (r UserRepo) Save(u *models.User) error {
	if u.ID != 0 {
		// User exists in db, should update
		_, err := r.DB.Exec(`
			UPDATE users
			SET
			    password_hash = $1,
			    name = $2,
			    email = $3
			WHERE
			    id = $4;
		`, u.Password, u.Name, u.Email, u.ID)
		if err != nil {
			return err
		}
	} else {
		// User doesn't exists in db, should insert
		err := r.DB.QueryRow(`
			INSERT INTO
			    users (username, password_hash, name, email)
			VALUES
			    ($1, $2, $3, $4)
			RETURNING
			    id;
		`, u.Username, u.Password, u.Name, u.Email).Scan(&u.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func getUserFromQuery(row *sql.Row) (*models.User, error) {
	user := models.User{}
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r UserRepo) GetByID(id int) (*models.User, error) {
	row := r.DB.QueryRow(`
		SELECT
		    id,
		    username,
		    password_hash,
		    name,
		    email
		FROM
		    users
		WHERE
		    id = $1;
	`, id)

	return getUserFromQuery(row)
}

func (r UserRepo) GetByUsername(username string) (*models.User, error) {
	row := r.DB.QueryRow(`
		SELECT
		    id,
		    username,
		    password_hash,
		    name,
		    email
		FROM
		    users
		WHERE
		    username = $1;
	`, username)

	return getUserFromQuery(row)
}

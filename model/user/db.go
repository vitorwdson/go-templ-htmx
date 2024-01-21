package user

import "database/sql"

func (u *User) Save(db *sql.DB) error {
	if u.ID != 0 {
		// User exists in db, should update
		if u.password != "" {
			_, err := db.Exec(`
                UPDATE
                    users
                SET
                    password = $1,
                    name = $2,
                    email = $3
                WHERE
                    id = $4;
            `, u.password, u.Name, u.Email, u.ID)
			if err != nil {
				return err
			}
		} else {
			_, err := db.Exec(`
                UPDATE
                    users
                SET
                    name = $1,
                    email = $2
                WHERE
                    id = $3;
            `, u.Name, u.Email, u.ID)
			if err != nil {
				return err
			}
		}
	} else {
		// User doesn't exists in db, should insert
		err := db.QueryRow(`
            INSERT INTO
                users (
                    username,
                    password,
                    name,
                    email
                )
            VALUES (
                $1,
                $2,
                $3,
                $4
            ) RETURNING id;
        `, u.Username, u.password, u.Name, u.Email).Scan(&u.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func getUserFromQuery(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.password,
		&user.Name,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetByID(db *sql.DB, id int) (*User, error) {
	row := db.QueryRow(`
        SELECT
            id,
            username,
            password,
            name,
            email
        FROM
            users
        WHERE
            id = $1;            
    `, id)

	return getUserFromQuery(row)
}

func GetByUsername(db *sql.DB, username string) (*User, error) {
	row := db.QueryRow(`
        SELECT
            id,
            username,
            password,
            name,
            email
        FROM
            users
        WHERE
            username = $1;            
    `, username)

	return getUserFromQuery(row)
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (username, password_hash, name, email)
VALUES
    ($1, $2, $3, $4)
RETURNING
    id, username, password_hash, name, email
`

type CreateUserParams struct {
	Username     string
	PasswordHash string
	Name         string
	Email        sql.NullString
}

// CreateUser
//
//	INSERT INTO
//	    users (username, password_hash, name, email)
//	VALUES
//	    ($1, $2, $3, $4)
//	RETURNING
//	    id, username, password_hash, name, email
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.PasswordHash,
		arg.Name,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.Name,
		&i.Email,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT
    id, username, password_hash, name, email
FROM
    users
WHERE
    id = $1
`

// GetUserByID
//
//	SELECT
//	    id, username, password_hash, name, email
//	FROM
//	    users
//	WHERE
//	    id = $1
func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.Name,
		&i.Email,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT
    id, username, password_hash, name, email
FROM
    users
WHERE
    username = $1
`

// GetUserByUsername
//
//	SELECT
//	    id, username, password_hash, name, email
//	FROM
//	    users
//	WHERE
//	    username = $1
func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.Name,
		&i.Email,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
    password_hash = $1,
    name = $2,
    email = $3
WHERE
    id = $4
`

type UpdateUserParams struct {
	PasswordHash string
	Name         string
	Email        sql.NullString
	ID           int32
}

// UpdateUser
//
//	UPDATE users
//	SET
//	    password_hash = $1,
//	    name = $2,
//	    email = $3
//	WHERE
//	    id = $4
func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.PasswordHash,
		arg.Name,
		arg.Email,
		arg.ID,
	)
	return err
}

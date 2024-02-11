-- name: UpdateUser :exec
UPDATE users
SET
    password_hash = $1,
    name = $2,
    email = $3
WHERE
    id = $4;


-- name: CreateUser :one
INSERT INTO
    users (username, password_hash, name, email)
VALUES
    ($1, $2, $3, $4)
RETURNING
    id;


-- name: GetUserByUsername :one
SELECT
    *
FROM
    users
WHERE
    username = $1;


-- name: GetUserByID :one
SELECT
    *
FROM
    users
WHERE
    id = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (name, email, username, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1 RETURNING *;
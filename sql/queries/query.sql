-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1 RETURNING *;
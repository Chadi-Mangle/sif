-- name: GetUserByName :one 
SELECT * FROM users WHERE first_name = $1 AND last_name = $2;

-- name: GetUserPasswordByName :one 
SELECT hash_password FROM users WHERE first_name = $1 AND last_name = $2;

-- name: ListUsers :many
SELECT * FROM users ORDER BY last_name, first_name;

-- name: CreateUser :one
INSERT INTO users (first_name, last_name, hash_password, is_activated, has_paid, is_admin) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: SetUserReservations :one
UPDATE users SET bungalow_id = $1 WHERE id = $2
RETURNING *;

-- name: SetUserPaid :one
UPDATE users SET has_paid = $1 WHERE id = $2
RETURNING *;

-- name: SetUserPassword :one
UPDATE users SET hash_password = $1 WHERE id = $2
RETURNING *;
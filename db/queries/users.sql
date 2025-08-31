-- name: GetUserByName :one 
SELECT * FROM users WHERE first_name = $1 AND last_name = $2;

-- name: ListUsers :many
SELECT * FROM users ORDER BY last_name, first_name;

-- name: CreateUser :one
INSERT INTO users (first_name, last_name, has_paid, bungalow_id) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: SetUserReservations :one
UPDATE users SET bungalow_id = $1 WHERE id = $2
RETURNING *;

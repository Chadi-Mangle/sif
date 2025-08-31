-- name: GetBungalowByID :one
SELECT * FROM bungalows_users WHERE id = $1;

-- name: ListBungalows :many
SELECT * FROM bungalows_users;

-- name: CreateBungalow :one
INSERT INTO bungalows (capacity) 
VALUES ($1)
RETURNING *;

-- name: GetBungalowNbReservations :one
SELECT COUNT(users.id) FROM users WHERE users.bungalow_id = $1;

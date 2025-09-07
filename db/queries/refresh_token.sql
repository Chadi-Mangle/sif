-- name: CreateRefreshToken :one
INSERT INTO refresh_token (id, first_name, last_name, token, expires_at) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM refresh_token WHERE id = $1 AND is_revoked = false;

-- name: GetRefreshTokenByToken :one
SELECT * FROM refresh_token WHERE token = $1 AND is_revoked = false;

-- name: RevokeRefreshToken :exec
UPDATE refresh_token SET is_revoked = true WHERE id = $1;

-- name: RevokeAllUserTokens :exec
UPDATE refresh_token SET is_revoked = true WHERE first_name = $1 AND last_name = $2;

-- name: DeleteExpiredTokens :exec
DELETE FROM refresh_token WHERE expires_at < NOW();
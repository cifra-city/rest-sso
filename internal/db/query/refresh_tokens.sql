-- name: InsertRefreshToken :exec
INSERT INTO refresh_tokens (id, user_id, token, created_at, expires_at, device_info, ip_address)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: GetRefreshToken :one
SELECT id, user_id, token, created_at, expires_at, device_info, ip_address
FROM refresh_tokens
WHERE token = $1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token = $1;

-- name: DeleteAllTokensForUser :exec
DELETE FROM refresh_tokens
WHERE user_id = $1;

-- name: UpdateRefreshToken :exec
UPDATE refresh_tokens
SET token = $2, expires_at = $3, device_info = $4, ip_address = $5
WHERE id = $1;

-- name: IsTokenExpired :one
SELECT COUNT(*)
FROM refresh_tokens
WHERE token = $1 AND expires_at > now();

-- name: GetTokensByUserID :many
SELECT id, user_id, token, created_at, expires_at, device_info, ip_address
FROM refresh_tokens
WHERE user_id = $1;

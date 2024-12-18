-- name: GetUserByID :one
SELECT * FROM users_secret
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users_secret
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users_secret
WHERE email = $1 LIMIT 1;

-- name: ListUsersByID :many
SELECT * FROM users_secret
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: ListUsersByUsername :many
SELECT * FROM users_secret
ORDER BY username
    LIMIT $1
OFFSET $2;

-- name: UpdateUsernameByID :one
UPDATE users_secret
SET username = $2
WHERE id = $1
    RETURNING *;


-- name: UpdateUserPasswordByID :one
UPDATE users_secret
SET pass_hash = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateEmailByID :one
UPDATE users_secret
SET
    email = $2
WHERE id = $1
    RETURNING *;

-- name: InsertUser :one
INSERT INTO users_secret (
    id,
    username,
    email,
    pass_hash
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: DeleteUserByID :exec
DELETE FROM users_secret
WHERE id = $1;

-- name: SetUserTokenVersionByID :one
UPDATE users_secret
SET token_version = $2
WHERE id = $1
    RETURNING *;

-- name: UpdateUserTokenVersionByID :one
UPDATE users_secret
SET token_version = token_version + 1
WHERE id = $1
    RETURNING *;
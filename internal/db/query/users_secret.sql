-- name: InsertUser :one
INSERT INTO users_secret (
    id,
    username,
    email,
    pass_hash
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users_secret
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users_secret
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users_secret
WHERE email = $1 LIMIT 1;

-- name: UpdateUsernameByID :one
UPDATE users_secret
SET
    username = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;


-- name: UpdateUserPasswordByID :one
UPDATE users_secret
SET
    pass_hash = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateEmailByID :one
UPDATE users_secret
SET
    email = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateUserTokenVersionByID :one
UPDATE users_secret
SET
    token_version = token_version + 1,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: SetUserTokenVersionByID :one
UPDATE users_secret
SET
    token_version = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: DeleteUserByID :exec
DELETE FROM users_secret
WHERE id = $1;


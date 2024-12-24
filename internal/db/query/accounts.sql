-- name: CreateAccount :one
INSERT INTO accounts (
    username,
    email,
    pass_hash
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByUsername :one
SELECT * FROM accounts
WHERE username = $1 LIMIT 1;

-- name: GetAccountByEmail :one
SELECT * FROM accounts
WHERE email = $1 LIMIT 1;

-- name: GetAccountByLogin :one
SELECT * FROM accounts
WHERE username = $1 AND email = $2 LIMIT 1;

-- name: AccountExists :one
SELECT EXISTS (
    SELECT 1 FROM accounts
    WHERE username = $1 OR email = $2
) AS account_exists;

-- name: UpdateUsername :one
UPDATE accounts
SET
    username = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdatePassword :one
UPDATE accounts
SET
    pass_hash = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;

-- name: UpdateTokenVersion :one
UPDATE accounts
SET
    token_version = token_version + 1,
    updated_at = now()
WHERE id = $1
    RETURNING *;
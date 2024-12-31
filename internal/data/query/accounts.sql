-- name: CreateAccount :one
INSERT INTO accounts (
    email,
    pass_hash
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetAccountByID :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByEmail :one
SELECT * FROM accounts
WHERE email = $1 LIMIT 1;

-- name: UpdatePassword :one
UPDATE accounts
SET
    pass_hash = $2,
    updated_at = now()
WHERE id = $1
    RETURNING *;
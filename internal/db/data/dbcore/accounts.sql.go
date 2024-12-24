// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package dbcore

import (
	"context"

	"github.com/google/uuid"
)

const accountExists = `-- name: AccountExists :one
SELECT EXISTS (
    SELECT 1 FROM accounts
    WHERE username = $1 OR email = $2
) AS account_exists
`

type AccountExistsParams struct {
	Username string
	Email    string
}

func (q *Queries) AccountExists(ctx context.Context, arg AccountExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, accountExists, arg.Username, arg.Email)
	var account_exists bool
	err := row.Scan(&account_exists)
	return account_exists, err
}

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    username,
    email,
    pass_hash
) VALUES (
    $1, $2, $3
) RETURNING id, username, email, pass_hash, role, token_version, created_at, updated_at
`

type CreateAccountParams struct {
	Username string
	Email    string
	PassHash string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Username, arg.Email, arg.PassHash)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.TokenVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT id, username, email, pass_hash, role, token_version, created_at, updated_at FROM accounts
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByEmail, email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.TokenVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByID = `-- name: GetAccountByID :one
SELECT id, username, email, pass_hash, role, token_version, created_at, updated_at FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccountByID(ctx context.Context, id uuid.UUID) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.TokenVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByLogin = `-- name: GetAccountByLogin :one
SELECT id, username, email, pass_hash, role, token_version, created_at, updated_at FROM accounts
WHERE username = $1 AND email = $2 LIMIT 1
`

type GetAccountByLoginParams struct {
	Username string
	Email    string
}

func (q *Queries) GetAccountByLogin(ctx context.Context, arg GetAccountByLoginParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByLogin, arg.Username, arg.Email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.TokenVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByUsername = `-- name: GetAccountByUsername :one
SELECT id, username, email, pass_hash, role, token_version, created_at, updated_at FROM accounts
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetAccountByUsername(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUsername, username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.TokenVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :one
UPDATE accounts
SET
    pass_hash = $2,
    updated_at = now()
WHERE id = $1
    RETURNING id, username, email, pass_hash, role, token_version, created_at, updated_at
`

type UpdatePasswordParams struct {
	ID       uuid.UUID
	PassHash string
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updatePassword, arg.ID, arg.PassHash)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.TokenVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTokenVersion = `-- name: UpdateTokenVersion :one
UPDATE accounts
SET
    token_version = token_version + 1,
    updated_at = now()
WHERE id = $1
    RETURNING id, username, email, pass_hash, role, token_version, created_at, updated_at
`

func (q *Queries) UpdateTokenVersion(ctx context.Context, id uuid.UUID) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateTokenVersion, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.TokenVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUsername = `-- name: UpdateUsername :one
UPDATE accounts
SET
    username = $2,
    updated_at = now()
WHERE id = $1
    RETURNING id, username, email, pass_hash, role, token_version, created_at, updated_at
`

type UpdateUsernameParams struct {
	ID       uuid.UUID
	Username string
}

func (q *Queries) UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateUsername, arg.ID, arg.Username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.TokenVersion,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package sqlcore

import (
	"context"

	"github.com/google/uuid"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    email,
    pass_hash
) VALUES (
    $1, $2
) RETURNING id, email, pass_hash, role, created_at, updated_at
`

type CreateAccountParams struct {
	Email    string
	PassHash string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Email, arg.PassHash)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT id, email, pass_hash, role, created_at, updated_at FROM accounts
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByEmail, email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByID = `-- name: GetAccountByID :one
SELECT id, email, pass_hash, role, created_at, updated_at FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccountByID(ctx context.Context, id uuid.UUID) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PassHash,
		&i.Role,
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
    RETURNING id, email, pass_hash, role, created_at, updated_at
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
		&i.Email,
		&i.PassHash,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

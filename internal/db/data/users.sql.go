// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package data

import (
	"context"

	"github.com/google/uuid"
)

const deleteUserByID = `-- name: DeleteUserByID :exec
DELETE FROM users_secret
WHERE id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserByID, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, email_status, pass_hash FROM users_secret
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, email_status, pass_hash FROM users_secret
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, email_status, pass_hash FROM users_secret
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users_secret (
    id,
    username,
    email,
    email_status,
    pass_hash
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, username, email, email_status, pass_hash
`

type InsertUserParams struct {
	ID          uuid.UUID
	Username    string
	Email       string
	EmailStatus bool
	PassHash    string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.EmailStatus,
		arg.PassHash,
	)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

const listUsersByID = `-- name: ListUsersByID :many
SELECT id, username, email, email_status, pass_hash FROM users_secret
ORDER BY id
    LIMIT $1
OFFSET $2
`

type ListUsersByIDParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsersByID(ctx context.Context, arg ListUsersByIDParams) ([]UsersSecret, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersSecret
	for rows.Next() {
		var i UsersSecret
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.EmailStatus,
			&i.PassHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsersByUsername = `-- name: ListUsersByUsername :many
SELECT id, username, email, email_status, pass_hash FROM users_secret
ORDER BY username
    LIMIT $1
OFFSET $2
`

type ListUsersByUsernameParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListUsersByUsername(ctx context.Context, arg ListUsersByUsernameParams) ([]UsersSecret, error) {
	rows, err := q.db.QueryContext(ctx, listUsersByUsername, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersSecret
	for rows.Next() {
		var i UsersSecret
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.EmailStatus,
			&i.PassHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEmailByID = `-- name: UpdateEmailByID :one
UPDATE users_secret
SET
    email = $2,
    email_status = $3
WHERE id = $1
    RETURNING id, username, email, email_status, pass_hash
`

type UpdateEmailByIDParams struct {
	ID          uuid.UUID
	Email       string
	EmailStatus bool
}

func (q *Queries) UpdateEmailByID(ctx context.Context, arg UpdateEmailByIDParams) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, updateEmailByID, arg.ID, arg.Email, arg.EmailStatus)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

const updateEmailStatusByID = `-- name: UpdateEmailStatusByID :one
UPDATE users_secret
SET email_status = $2
WHERE id = $1
    RETURNING id, username, email, email_status, pass_hash
`

type UpdateEmailStatusByIDParams struct {
	ID          uuid.UUID
	EmailStatus bool
}

func (q *Queries) UpdateEmailStatusByID(ctx context.Context, arg UpdateEmailStatusByIDParams) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, updateEmailStatusByID, arg.ID, arg.EmailStatus)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

const updatePasswordByID = `-- name: UpdatePasswordByID :one
UPDATE users_secret
SET pass_hash = $2
WHERE id = $1
    RETURNING id, username, email, email_status, pass_hash
`

type UpdatePasswordByIDParams struct {
	ID       uuid.UUID
	PassHash string
}

func (q *Queries) UpdatePasswordByID(ctx context.Context, arg UpdatePasswordByIDParams) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, updatePasswordByID, arg.ID, arg.PassHash)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

const updateUserByID = `-- name: UpdateUserByID :one
UPDATE users_secret
SET
    email = $2,
    email_status = $3,
    username = $4
WHERE id = $1
    RETURNING id, username, email, email_status, pass_hash
`

type UpdateUserByIDParams struct {
	ID          uuid.UUID
	Email       string
	EmailStatus bool
	Username    string
}

func (q *Queries) UpdateUserByID(ctx context.Context, arg UpdateUserByIDParams) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, updateUserByID,
		arg.ID,
		arg.Email,
		arg.EmailStatus,
		arg.Username,
	)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

const updateUsernameByID = `-- name: UpdateUsernameByID :one
UPDATE users_secret
SET username = $2
WHERE id = $1
    RETURNING id, username, email, email_status, pass_hash
`

type UpdateUsernameByIDParams struct {
	ID       uuid.UUID
	Username string
}

func (q *Queries) UpdateUsernameByID(ctx context.Context, arg UpdateUsernameByIDParams) (UsersSecret, error) {
	row := q.db.QueryRowContext(ctx, updateUsernameByID, arg.ID, arg.Username)
	var i UsersSecret
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.EmailStatus,
		&i.PassHash,
	)
	return i, err
}

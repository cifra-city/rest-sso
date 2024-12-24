// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sessions.sql

package dbcore

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :exec
INSERT INTO sessions (user_id, token, device_data)
VALUES ($1, $2, $3)
`

type CreateSessionParams struct {
	UserID     uuid.UUID
	Token      string
	DeviceData json.RawMessage
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := q.db.ExecContext(ctx, createSession, arg.UserID, arg.Token, arg.DeviceData)
	return err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSession, id)
	return err
}

const deleteUserSession = `-- name: DeleteUserSession :exec
DELETE FROM sessions
WHERE id = $1 AND user_id = $2
`

type DeleteUserSessionParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) DeleteUserSession(ctx context.Context, arg DeleteUserSessionParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserSession, arg.ID, arg.UserID)
	return err
}

const deleteUserSessions = `-- name: DeleteUserSessions :exec
DELETE FROM sessions
WHERE user_id = $1
`

func (q *Queries) DeleteUserSessions(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserSessions, userID)
	return err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, token, device_data, created_at, last_used FROM sessions
WHERE id = $1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.DeviceData,
		&i.CreatedAt,
		&i.LastUsed,
	)
	return i, err
}

const getSessionToken = `-- name: GetSessionToken :one
SELECT token FROM sessions
WHERE id = $1 AND user_id = $2
`

type GetSessionTokenParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) GetSessionToken(ctx context.Context, arg GetSessionTokenParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getSessionToken, arg.ID, arg.UserID)
	var token string
	err := row.Scan(&token)
	return token, err
}

const getSessionsByUserID = `-- name: GetSessionsByUserID :many
SELECT id, user_id, token, device_data, created_at, last_used FROM sessions
WHERE user_id = $1
`

func (q *Queries) GetSessionsByUserID(ctx context.Context, userID uuid.UUID) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, getSessionsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Token,
			&i.DeviceData,
			&i.CreatedAt,
			&i.LastUsed,
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

const getUserSession = `-- name: GetUserSession :one
SELECT id, user_id, token, device_data, created_at, last_used FROM sessions
WHERE id = $1 AND user_id = $2
`

type GetUserSessionParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) GetUserSession(ctx context.Context, arg GetUserSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, getUserSession, arg.ID, arg.UserID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.DeviceData,
		&i.CreatedAt,
		&i.LastUsed,
	)
	return i, err
}

const updateSessionToken = `-- name: UpdateSessionToken :exec
UPDATE sessions
SET
    token = $3,
    last_used = now()
WHERE id = $1 AND user_id = $2
`

type UpdateSessionTokenParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Token  string
}

func (q *Queries) UpdateSessionToken(ctx context.Context, arg UpdateSessionTokenParams) error {
	_, err := q.db.ExecContext(ctx, updateSessionToken, arg.ID, arg.UserID, arg.Token)
	return err
}

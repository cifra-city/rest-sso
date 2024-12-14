// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: refresh_tokens.sql

package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const deleteAllTokensForUser = `-- name: DeleteAllTokensForUser :exec
DELETE FROM refresh_tokens
WHERE user_id = $1
`

func (q *Queries) DeleteAllTokensForUser(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllTokensForUser, userID)
	return err
}

const deleteRefreshToken = `-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token = $1
`

func (q *Queries) DeleteRefreshToken(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deleteRefreshToken, token)
	return err
}

const getRefreshToken = `-- name: GetRefreshToken :one
SELECT id, user_id, token, created_at, expires_at, device_id, ip_address
FROM refresh_tokens
WHERE token = $1
`

func (q *Queries) GetRefreshToken(ctx context.Context, token string) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, getRefreshToken, token)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.DeviceID,
		&i.IpAddress,
	)
	return i, err
}

const getTokenByUserIdAndDeviceId = `-- name: GetTokenByUserIdAndDeviceId :one
SELECT id, user_id, token, created_at, expires_at, device_id, ip_address
FROM refresh_tokens
WHERE user_id = $1 AND device_id = $2
`

type GetTokenByUserIdAndDeviceIdParams struct {
	UserID   uuid.UUID
	DeviceID uuid.UUID
}

func (q *Queries) GetTokenByUserIdAndDeviceId(ctx context.Context, arg GetTokenByUserIdAndDeviceIdParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, getTokenByUserIdAndDeviceId, arg.UserID, arg.DeviceID)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.CreatedAt,
		&i.ExpiresAt,
		&i.DeviceID,
		&i.IpAddress,
	)
	return i, err
}

const getTokensByUserID = `-- name: GetTokensByUserID :many
SELECT id, user_id, token, created_at, expires_at, device_id, ip_address
FROM refresh_tokens
WHERE user_id = $1
`

func (q *Queries) GetTokensByUserID(ctx context.Context, userID uuid.UUID) ([]RefreshToken, error) {
	rows, err := q.db.QueryContext(ctx, getTokensByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RefreshToken
	for rows.Next() {
		var i RefreshToken
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Token,
			&i.CreatedAt,
			&i.ExpiresAt,
			&i.DeviceID,
			&i.IpAddress,
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

const insertRefreshToken = `-- name: InsertRefreshToken :exec
INSERT INTO refresh_tokens (id, user_id, token, created_at, expires_at, device_id, ip_address)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type InsertRefreshTokenParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	CreatedAt time.Time
	ExpiresAt time.Time
	DeviceID  uuid.UUID
	IpAddress sql.NullString
}

func (q *Queries) InsertRefreshToken(ctx context.Context, arg InsertRefreshTokenParams) error {
	_, err := q.db.ExecContext(ctx, insertRefreshToken,
		arg.ID,
		arg.UserID,
		arg.Token,
		arg.CreatedAt,
		arg.ExpiresAt,
		arg.DeviceID,
		arg.IpAddress,
	)
	return err
}

const isTokenExpired = `-- name: IsTokenExpired :one
SELECT COUNT(*)
FROM refresh_tokens
WHERE token = $1 AND expires_at > now()
`

func (q *Queries) IsTokenExpired(ctx context.Context, token string) (int64, error) {
	row := q.db.QueryRowContext(ctx, isTokenExpired, token)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateRefreshToken = `-- name: UpdateRefreshToken :exec
UPDATE refresh_tokens
SET token = $2, expires_at = $3, device_id = $4, ip_address = $5
WHERE id = $1
`

type UpdateRefreshTokenParams struct {
	ID        uuid.UUID
	Token     string
	ExpiresAt time.Time
	DeviceID  uuid.UUID
	IpAddress sql.NullString
}

func (q *Queries) UpdateRefreshToken(ctx context.Context, arg UpdateRefreshTokenParams) error {
	_, err := q.db.ExecContext(ctx, updateRefreshToken,
		arg.ID,
		arg.Token,
		arg.ExpiresAt,
		arg.DeviceID,
		arg.IpAddress,
	)
	return err
}

const updateRefreshTokenByDeviceAndUserID = `-- name: UpdateRefreshTokenByDeviceAndUserID :exec
UPDATE refresh_tokens
SET token = $3, expires_at = $4
WHERE user_id = $1 AND device_id = $2
`

type UpdateRefreshTokenByDeviceAndUserIDParams struct {
	UserID    uuid.UUID
	DeviceID  uuid.UUID
	Token     string
	ExpiresAt time.Time
}

func (q *Queries) UpdateRefreshTokenByDeviceAndUserID(ctx context.Context, arg UpdateRefreshTokenByDeviceAndUserIDParams) error {
	_, err := q.db.ExecContext(ctx, updateRefreshTokenByDeviceAndUserID,
		arg.UserID,
		arg.DeviceID,
		arg.Token,
		arg.ExpiresAt,
	)
	return err
}

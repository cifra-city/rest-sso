package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// UpdateRefreshTokenTransaction updates the refresh token for a user and device,
// creating 2 new records if necessary: a new device, a new refresh token.
// or updating refresh_token if device its first usage creating new record in device, always
func (q *Queries) UpdateRefreshTokenTransaction(
	ctx context.Context,
	userID uuid.UUID,
	deviceID uuid.UUID,
	factoryID string,
	deviceName string,
	osVersion string,
	newToken string,
	expiresAt time.Time,
	ipAddress string) error {

	tx, err := q.db.(*sql.DB).BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	queries := q.WithTx(tx)

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = queries.UpdateLastUsed(ctx, UpdateLastUsedParams{
		ID:       deviceID,
		LastUsed: time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	_, err = queries.GetTokenByUserIdAndDeviceId(ctx, GetTokenByUserIdAndDeviceIdParams{
		UserID:   userID,
		DeviceID: deviceID,
	})
	if err != nil {
		return err
	}

	err = queries.UpdateRefreshTokenByDeviceAndUserID(ctx, UpdateRefreshTokenByDeviceAndUserIDParams{
		UserID:    userID,
		DeviceID:  deviceID,
		Token:     newToken,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

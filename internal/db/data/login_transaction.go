package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// LoginTransaction updates the refresh token for a user and device,
// creating 2 new records if necessary: a new device, a new refresh token.
// or updating refresh_token if device its first usage creating new record in device, always
func (q *Queries) LoginTransaction(
	ctx context.Context,
	userID uuid.UUID,
	deviceID uuid.UUID,
	encryptedToken string,
	expiresAt time.Time,
	factoryID string,
	deviceName string,
	osVersion string,
	ipAddress string,
	fingerprint string) error {

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

	err = queries.InsertDevice(ctx, InsertDeviceParams{
		ID:         deviceID,
		UserID:     userID,
		FactoryID:  factoryID,
		DeviceName: sql.NullString{String: deviceName, Valid: true},
		OsVersion:  sql.NullString{String: osVersion, Valid: true},
		CreatedAt:  time.Now().UTC(),
		LastUsed:   time.Now().UTC(),
	})
	if err != nil {
		return err
	}

	device, err := queries.GetDeviceByID(ctx, deviceID)
	if err != nil {
		return err
	}

	err = queries.InsertRefreshToken(ctx, InsertRefreshTokenParams{
		ID:        uuid.New(),
		UserID:    userID,
		Token:     encryptedToken,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: expiresAt,
		DeviceID:  device.ID,
		IpAddress: ipAddress,
	})
	if err != nil {
		return err
	}

	err = queries.InsertOperationHistory(ctx, InsertOperationHistoryParams{
		ID:         uuid.New(),
		UserID:     userID,
		DeviceData: fingerprint,
		Operation:  OperationTypeLogin,
		Success:    true,
		IpAddress:  ipAddress,
	})

	return tx.Commit()
}

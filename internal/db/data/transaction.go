package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

// UpdateRefreshTokenTransaction updates the refresh token for a user and device,
// creating 3 new records if necessary: a new device, a new refresh token, and a new login history record in db.
// or updating refresh_token if device its first usage creating new record in device, always create new login_history
func (q *Queries) UpdateRefreshTokenTransaction(
	ctx context.Context,
	user *UsersSecret,
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

	device, err := queries.GetDeviceByUserIDAndFactoryId(ctx, GetDeviceByUserIDAndFactoryIdParams{
		UserID:    user.ID,
		FactoryID: factoryID,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if errors.Is(err, sql.ErrNoRows) { // Device is nonexistent
		err := queries.InsertDevice(ctx, InsertDeviceParams{
			ID:         uuid.New(),
			UserID:     user.ID,
			FactoryID:  factoryID,
			DeviceName: sql.NullString{String: deviceName, Valid: true},
			OsVersion:  sql.NullString{String: osVersion, Valid: true},
			CreatedAt:  time.Now().UTC(),
			LastUsed:   time.Now().UTC(),
		})
		if err != nil {
			return err
		}

		device, err = queries.GetDeviceByUserIDAndFactoryId(ctx, GetDeviceByUserIDAndFactoryIdParams{
			UserID:    user.ID,
			FactoryID: factoryID,
		})
		if err != nil {
			return err
		}
	} else { // if device exists
		err := queries.UpdateLastUsed(ctx, UpdateLastUsedParams{
			ID:       device.ID,
			LastUsed: time.Now().UTC(),
		})
		if err != nil {
			return err
		}
	}

	_, err = queries.GetTokenByUserIdAndDeviceId(ctx, GetTokenByUserIdAndDeviceIdParams{
		UserID:   user.ID,
		DeviceID: device.ID,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if errors.Is(err, sql.ErrNoRows) { // Token for this user isn`t exists
		err := queries.InsertRefreshToken(ctx, InsertRefreshTokenParams{
			ID:        uuid.New(),
			UserID:    user.ID,
			Token:     newToken,
			CreatedAt: time.Now().UTC(),
			ExpiresAt: expiresAt,
			DeviceID:  device.ID,
			IpAddress: ipAddress,
		})
		if err != nil {
			return err
		}
	} else if err == nil { // Token is exists
		err = queries.UpdateRefreshTokenByDeviceAndUserID(ctx, UpdateRefreshTokenByDeviceAndUserIDParams{
			UserID:    user.ID,
			DeviceID:  device.ID,
			Token:     newToken,
			ExpiresAt: expiresAt,
		})
		if err != nil {
			return err
		}
	} else {
		return err
	}

	err = queries.InsertLoginHistory(ctx, InsertLoginHistoryParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		DeviceID:  device.ID,
		IpAddress: ipAddress,
		LoginTime: time.Now().UTC(),
		Success:   true,
		FailureReason: NullFailureReason{
			FailureReason: FailureReasonSuccess,
			Valid:         true,
		},
	})
	return tx.Commit()
}

// ResetPassword updates the refresh token for a user and device,
func (q *Queries) ResetPassword(
	ctx context.Context,
	user *UsersSecret,
	newToken string,
	expiresAt time.Time,
	hashedPassword string) error {

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

	devices, err := queries.GetDevicesByUserID(ctx, user.ID)
	if err != nil {
		return err
	}

	for _, device := range devices {
		err = queries.UpdateRefreshTokenByDeviceAndUserID(ctx, UpdateRefreshTokenByDeviceAndUserIDParams{
			UserID:    user.ID,
			DeviceID:  device.ID,
			Token:     newToken,
			ExpiresAt: expiresAt,
		})
		if err != nil {
			return err
		}
	}

	_, err = queries.UpdateUserPasswordByID(ctx, UpdateUserPasswordByIDParams{
		ID:       user.ID,
		PassHash: hashedPassword,
	})
	if err != nil {
		return err
	}

	_, err = queries.UpdateUserTokenVersionByID(ctx, user.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

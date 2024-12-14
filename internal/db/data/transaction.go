package data

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
		log.Fatalf("failed to begin transaction: %v", err)
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
	if err != nil && !errors.Is(err, sql.ErrNoRows) { // Internal Error
		logrus.Errorf("err device 1: %s", err)
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
			logrus.Errorf("err device 2: %s", err)
			return err
		}

		device, err = queries.GetDeviceByUserIDAndFactoryId(ctx, GetDeviceByUserIDAndFactoryIdParams{
			UserID:    user.ID,
			FactoryID: factoryID,
		})
		if err != nil {
			logrus.Errorf("err device 3: %s", err)
			return err
		}
	} else { // if device exists
		err := queries.UpdateLastUsed(ctx, UpdateLastUsedParams{
			ID:       device.ID,
			LastUsed: time.Now().UTC(),
		})
		if err != nil {
			logrus.Errorf("err device 4: %s", err)
			return err
		}
	}

	_, err = queries.GetTokenByUserIdAndDeviceId(ctx, GetTokenByUserIdAndDeviceIdParams{
		UserID:   user.ID,
		DeviceID: device.ID,
	})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logrus.Errorf("err token 1: %s", err)
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
			IpAddress: sql.NullString{String: ipAddress, Valid: true},
		})
		if err != nil {
			logrus.Errorf("err token 2: %s", err)
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
			logrus.Errorf("err token 3: %s", err)
			return err
		}
	} else {
		logrus.Errorf("err token 4: %s", err)
		return err
	}

	err = queries.InsertLoginHistory(ctx, InsertLoginHistoryParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		DeviceID:  uuid.NullUUID{UUID: device.ID, Valid: true},
		IpAddress: sql.NullString{String: ipAddress, Valid: true},
		LoginTime: time.Now().UTC(),
		Success:   true,
		FailureReason: NullFailureReason{
			FailureReason: FailureReasonSuccess,
			Valid:         true,
		},
	})
	logrus.Errorf("err final 1: %s", err)
	return tx.Commit()
}

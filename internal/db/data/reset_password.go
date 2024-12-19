package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// ResetPasswordTransaction updates the refresh token for a user and device,
func (q *Queries) ResetPasswordTransaction(
	ctx context.Context,
	user *UsersSecret,
	newToken string,
	expiresAt time.Time,
	hashedPassword string,
	deviceData string,
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

	err = queries.InsertOperationHistory(ctx, InsertOperationHistoryParams{
		ID:            uuid.New(),
		UserID:        user.ID,
		DeviceData:    deviceData,
		Operation:     OperationTypeChangePassword,
		Success:       true,
		FailureReason: FailureReasonSuccess,
		IpAddress:     ipAddress,
	})

	return tx.Commit()
}

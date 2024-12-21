package data

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrorDeviceDoesNotBelongToUser = errors.New("device does not belong to user")
)

func (q *Queries) TerminateSessionsTransaction(ctx context.Context, devices []Device, userId uuid.UUID, fingerprint string, IP string) error {

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

	for _, dev := range devices {
		device, err := queries.GetDeviceByID(ctx, dev.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return err
			}
			return err
		}

		if device.UserID != userId {
			return ErrorDeviceDoesNotBelongToUser
		}

		err = queries.DeleteDeviceByUserIdAndId(ctx, DeleteDeviceByUserIdAndIdParams{
			UserID: userId,
			ID:     dev.ID,
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return err
			}
			return err
		}
	}

	err = queries.InsertOperationHistory(ctx, InsertOperationHistoryParams{
		ID:            uuid.New(),
		UserID:        userId,
		DeviceData:    fingerprint,
		Operation:     OperationTypeTerminateSession,
		Success:       true,
		FailureReason: FailureReasonSuccess,
		IpAddress:     IP,
	})

	return tx.Commit()
}

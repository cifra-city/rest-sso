package data

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

// ResetPasswordTransaction updates the refresh token for a user and device,
func (q *Queries) ResetPasswordTransaction(
	ctx context.Context,
	user *UsersSecret,
	hashedPassword string,
	fingerprint string,
	ipAddress string,
) error {

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

	err = queries.DeleteAllTokensForUser(ctx, user.ID)

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
		DeviceData:    fingerprint,
		Operation:     OperationTypeChangePassword,
		Success:       true,
		FailureReason: FailureReasonSuccess,
		IpAddress:     ipAddress,
	})

	return tx.Commit()
}

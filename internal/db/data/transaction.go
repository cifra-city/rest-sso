package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/google/uuid"
)

type Transaction interface {
	ResetPasswordTxn(
		ctx context.Context,
		userID uuid.UUID,
		hashedPassword string,
		deviceData json.RawMessage,
	) error

	LoginTxn(
		ctx context.Context,
		userID uuid.UUID,
		Token string,
		deviceData json.RawMessage,
	) (*dbcore.Session, error)

	TerminateSessionsTxn(
		ctx context.Context,
		userId uuid.UUID,
		curDevId uuid.UUID,
		deviceData json.RawMessage,
	) error

	UpdateRefreshTokenTrx( //TODO for future use right no sense
		ctx context.Context,
		userID uuid.UUID,
		sessionID uuid.UUID,
		newToken string,
		deviceData json.RawMessage) error
}

type transaction struct {
	queries *dbcore.Queries
}

func NewTransaction(queries *dbcore.Queries) Transaction {
	return &transaction{queries: queries}
}

func HandleTransactionRollback(tx *sql.Tx, originalErr error) error {
	if originalErr != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", originalErr, rbErr)
		}
	}
	return originalErr
}

func (t *transaction) ResetPasswordTxn(
	ctx context.Context,
	userID uuid.UUID,
	hashedPassword string,
	deviceData json.RawMessage,
) error {
	queries, tx, err := t.queries.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	_, err = queries.UpdatePassword(ctx, dbcore.UpdatePasswordParams{
		ID:       userID,
		PassHash: hashedPassword,
	})
	if err != nil {
		return err
	}

	_, err = queries.UpdateTokenVersion(ctx, userID)
	if err != nil {
		return err
	}

	err = queries.DeleteUserSessions(ctx, userID)
	if err != nil {
		return err
	}

	err = queries.CreateOperation(ctx, dbcore.CreateOperationParams{
		UserID:     userID,
		Operation:  dbcore.OperationTypeResetPassword,
		DeviceData: deviceData,
		Success:    true,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t *transaction) LoginTxn(
	ctx context.Context,
	userID uuid.UUID,
	Token string,
	deviceData json.RawMessage,
) (*dbcore.Session, error) {
	queries, tx, err := t.queries.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	session, err := queries.CreateSession(ctx, dbcore.CreateSessionParams{
		UserID:     userID,
		Token:      Token,
		DeviceData: deviceData,
	})
	if err != nil {
		return nil, err
	}

	err = queries.CreateOperation(ctx, dbcore.CreateOperationParams{
		UserID:     userID,
		Operation:  dbcore.OperationTypeLogin,
		DeviceData: deviceData,
		Success:    true,
	})
	if err != nil {
		return nil, err
	}

	return &session, tx.Commit()
}

func (t *transaction) TerminateSessionsTxn(
	ctx context.Context,
	userId uuid.UUID,
	curDevId uuid.UUID,
	deviceData json.RawMessage,
) error {
	queries, tx, err := t.queries.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	userSessions, err := queries.GetSessionsByUserID(ctx, userId)
	if err != nil {
		return err
	}

	for _, dev := range userSessions {
		if dev.ID == curDevId {
			continue
		}

		err = queries.DeleteUserSession(ctx, dbcore.DeleteUserSessionParams{
			ID:     dev.ID,
			UserID: userId,
		})
		if err != nil {
			return err
		}
	}

	err = queries.CreateOperation(ctx, dbcore.CreateOperationParams{
		UserID:     userId,
		Operation:  dbcore.OperationTypeTerminateSession,
		DeviceData: deviceData,
		Success:    true,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t *transaction) UpdateRefreshTokenTrx( //TODO for future use
	ctx context.Context,
	userID uuid.UUID,
	sessionID uuid.UUID,
	newToken string,
	deviceData json.RawMessage) error {

	queries, tx, err := t.queries.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	err = queries.UpdateSessionToken(ctx, dbcore.UpdateSessionTokenParams{
		ID:     sessionID,
		UserID: userID,
		Token:  newToken,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

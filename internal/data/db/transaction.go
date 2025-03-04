package db

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/rest-sso/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type Transaction interface {
	ResetPasswordTxn(
		r *http.Request,
		userID uuid.UUID,
		hashedPassword string,
	) error

	LoginTxn(
		r *http.Request,
		userID uuid.UUID,
		deviceName string,
		Token string,
		deviceID uuid.UUID,
	) (*sqlcore.Session, error)

	TerminateSessionsTxn(
		r *http.Request,
		userId uuid.UUID,
		curDevId uuid.UUID,
	) error

	UpdateRefreshTokenTrx( //TODO for future use right no sense
		r *http.Request,
		userID uuid.UUID,
		sessionID uuid.UUID,
		newToken string,
	) error
}

type transaction struct {
	queries *sqlcore.Queries
}

func NewTransaction(queries *sqlcore.Queries) Transaction {
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
	r *http.Request,
	userID uuid.UUID,
	hashedPassword string,
) error {
	ctx := r.Context()
	queries, tx, err := t.queries.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	deviceData, err := NewDeviceData(r)
	if err != nil {
		return err
	}

	_, err = queries.UpdatePassword(ctx, sqlcore.UpdatePasswordParams{
		ID:       userID,
		PassHash: hashedPassword,
	})
	if err != nil {
		return err
	}

	err = queries.DeleteUserSessions(ctx, userID)
	if err != nil {
		return err
	}

	err = queries.CreateOperation(ctx, sqlcore.CreateOperationParams{
		UserID:     userID,
		Operation:  sqlcore.OperationTypeResetPassword,
		DeviceData: deviceData,
		Success:    true,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t *transaction) LoginTxn(
	r *http.Request,
	userID uuid.UUID,
	deviceName string,
	Token string,
	deviceID uuid.UUID,
) (*sqlcore.Session, error) {
	ctx := r.Context()
	queries, tx, err := t.queries.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	client := httpkit.GetUserAgent(r)
	IP := httpkit.GetClientIP(r)

	deviceData, err := NewDeviceData(r)
	if err != nil {
		return nil, err
	}

	session, err := queries.CreateSession(ctx, sqlcore.CreateSessionParams{
		ID:         deviceID,
		UserID:     userID,
		Token:      Token,
		DeviceName: deviceName,
		Ip:         IP,
		Client:     client,
	})
	if err != nil {
		return nil, err
	}

	err = queries.CreateOperation(ctx, sqlcore.CreateOperationParams{
		UserID:     userID,
		Operation:  sqlcore.OperationTypeLogin,
		DeviceData: deviceData,
		Success:    true,
	})
	if err != nil {
		return nil, err
	}

	return &session, tx.Commit()
}

func (t *transaction) TerminateSessionsTxn(
	r *http.Request,
	userId uuid.UUID,
	curDevId uuid.UUID,
) error {
	ctx := r.Context()
	queries, tx, err := t.queries.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	deviceData, err := NewDeviceData(r)
	if err != nil {
		return err
	}

	userSessions, err := queries.GetSessionsByUserID(ctx, userId)
	if err != nil {
		return err
	}

	for _, dev := range userSessions {
		if dev.ID == curDevId {
			continue
		}
		err = queries.DeleteUserSession(ctx, sqlcore.DeleteUserSessionParams{
			ID:     dev.ID,
			UserID: userId,
		})
		if err != nil {
			return err
		}
	}

	err = queries.CreateOperation(ctx, sqlcore.CreateOperationParams{
		UserID:     userId,
		Operation:  sqlcore.OperationTypeTerminateSession,
		DeviceData: deviceData,
		Success:    true,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t *transaction) UpdateRefreshTokenTrx( //TODO for future use
	r *http.Request,
	userID uuid.UUID,
	sessionID uuid.UUID,
	newToken string) error {

	ctx := r.Context()
	queries, tx, err := t.queries.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	err = queries.UpdateSessionToken(ctx, sqlcore.UpdateSessionTokenParams{
		ID:     sessionID,
		UserID: userID,
		Token:  newToken,
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

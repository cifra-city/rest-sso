package data

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
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
	) (*dbcore.Session, error)

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
	r *http.Request,
	userID uuid.UUID,
	deviceName string,
	Token string,
) (*dbcore.Session, error) {
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

	session, err := queries.CreateSession(ctx, dbcore.CreateSessionParams{
		UserID:     userID,
		Token:      Token,
		Ip:         IP,
		Client:     client,
		DeviceName: deviceName,
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

	defer func() {
		err = HandleTransactionRollback(tx, err)
	}()

	userSessions, err := queries.GetSessionsByUserID(ctx, userId)
	if err != nil {
		return err
	}

	for _, dev := range userSessions {
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

	//deviceData, err := dbcore.NewDeviceData(r)
	//if err != nil {
	//	return err
	//}
	//if err != nil {
	//	return err
	//}

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

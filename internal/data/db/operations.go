package db

import (
	"net/http"

	"github.com/cifra-city/rest-sso/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type Operations interface {
	Create(r *http.Request, userID uuid.UUID, operation sqlcore.OperationType, success bool, failureReason *sqlcore.FailureReason) error
	CreateFailure(r *http.Request, userID uuid.UUID, operation sqlcore.OperationType, failureReason sqlcore.FailureReason) error
}
type operations struct {
	queries *sqlcore.Queries
}

func NewOperations(queries *sqlcore.Queries) Operations {
	return &operations{queries: queries}
}

func (o *operations) Create(r *http.Request, userID uuid.UUID, operation sqlcore.OperationType, success bool, failureReason *sqlcore.FailureReason) error {
	deviceData, err := NewDeviceData(r)
	if err != nil {
		return err
	}

	fr := sqlcore.NullFailureReason{
		FailureReason: "",
		Valid:         false,
	}

	if failureReason != nil && !success {
		fr = sqlcore.NullFailureReason{
			FailureReason: *failureReason,
			Valid:         true,
		}
	}

	return o.queries.CreateOperation(r.Context(), sqlcore.CreateOperationParams{
		UserID:        userID,
		Operation:     operation,
		DeviceData:    deviceData,
		Success:       success,
		FailureReason: fr,
	})
}

func (o *operations) CreateFailure(r *http.Request, userID uuid.UUID, operation sqlcore.OperationType, failureReason sqlcore.FailureReason) error {
	deviceData, err := NewDeviceData(r)
	if err != nil {
		return err
	}

	return o.queries.CreateOperation(r.Context(), sqlcore.CreateOperationParams{
		UserID:     userID,
		Operation:  operation,
		DeviceData: deviceData,
		Success:    false,
		FailureReason: sqlcore.NullFailureReason{
			FailureReason: failureReason,
			Valid:         true,
		},
	})
}

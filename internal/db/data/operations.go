package data

import (
	"net/http"

	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/google/uuid"
)

type Operations interface {
	Create(r *http.Request, userID uuid.UUID, operation dbcore.OperationType, success bool, failureReason *dbcore.FailureReason) error
	CreateFailure(r *http.Request, userID uuid.UUID, operation dbcore.OperationType, failureReason dbcore.FailureReason) error
}
type operations struct {
	queries *dbcore.Queries
}

func NewOperations(queries *dbcore.Queries) Operations {
	return &operations{queries: queries}
}

func (o *operations) Create(r *http.Request, userID uuid.UUID, operation dbcore.OperationType, success bool, failureReason *dbcore.FailureReason) error {
	deviceData, err := NewDeviceData(r)
	if err != nil {
		return err
	}

	fr := dbcore.NullFailureReason{
		FailureReason: "",
		Valid:         false,
	}

	if failureReason != nil && !success {
		fr = dbcore.NullFailureReason{
			FailureReason: *failureReason,
			Valid:         true,
		}
	}

	return o.queries.CreateOperation(r.Context(), dbcore.CreateOperationParams{
		UserID:        userID,
		Operation:     operation,
		DeviceData:    deviceData,
		Success:       success,
		FailureReason: fr,
	})
}

func (o *operations) CreateFailure(r *http.Request, userID uuid.UUID, operation dbcore.OperationType, failureReason dbcore.FailureReason) error {
	deviceData, err := NewDeviceData(r)
	if err != nil {
		return err
	}

	return o.queries.CreateOperation(r.Context(), dbcore.CreateOperationParams{
		UserID:     userID,
		Operation:  operation,
		DeviceData: deviceData,
		Success:    false,
		FailureReason: dbcore.NullFailureReason{
			FailureReason: failureReason,
			Valid:         true,
		},
	})
}

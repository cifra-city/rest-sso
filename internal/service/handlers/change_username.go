package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ChangeUsername(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewChangeUsername(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	oldPassword := req.Data.Attributes.Password
	newUsername := req.Data.Attributes.NewUsername

	IP := httpkit.GetClientIP(r)
	fingerprint := httpkit.GenerateFingerprint(r)

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	user, err := Server.Queries.GetUserByID(r.Context(), userID)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve user information"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(oldPassword))
	if err != nil {
		err = Server.Queries.InsertOperationHistory(r.Context(), data.InsertOperationHistoryParams{
			ID:            uuid.New(),
			UserID:        userID,
			DeviceData:    fingerprint,
			Operation:     data.OperationTypeChangeUsername,
			Success:       false,
			IpAddress:     IP,
			FailureReason: data.FailureReasonInvalidPassword,
		})
		if err != nil {
			log.Errorf("Failed to insert operation history: %v", err)
		}
		httpkit.RenderErr(w, problems.Unauthorized("Invalid password"))
		return
	}

	_, err = Server.Queries.GetUserByUsername(r.Context(), *newUsername)
	if !errors.Is(err, sql.ErrNoRows) {
		if err != nil {
			httpkit.RenderErr(w, problems.InternalError("Failed to check username availability"))
			return
		}
		httpkit.RenderErr(w, problems.Conflict("Username already exists"))
		return
	}

	_, err = Server.Queries.UpdateUsernameByID(r.Context(), data.UpdateUsernameByIDParams{
		ID:       userID,
		Username: *newUsername,
	})
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to update username"))
		return
	}

	err = Server.Queries.InsertOperationHistory(r.Context(), data.InsertOperationHistoryParams{
		ID:            uuid.New(),
		UserID:        userID,
		DeviceData:    fingerprint,
		Operation:     data.OperationTypeChangeUsername,
		Success:       true,
		IpAddress:     IP,
		FailureReason: data.FailureReasonSuccess,
	})
	if err != nil {
		log.Errorf("Failed to insert operation history: %v", err)
	}

	httpkit.Render(w, map[string]string{"message": "Username updated successfully"})
}

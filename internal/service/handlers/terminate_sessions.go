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
)

func TerminateSessions(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewTerminateSessions(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	devices := req.Data.Attributes.Devices

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

	log.Infof("userID: %v", userID)

	_, err = Server.Queries.GetUserByID(r.Context(), userID)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve user information"))
		return
	}

	var userSessions []data.Device
	for _, id := range devices {
		device, err := Server.Queries.GetDeviceByID(r.Context(), uuid.MustParse(id.Id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				httpkit.RenderErr(w, problems.NotFound("Device not found"))
				return
			}
			httpkit.RenderErr(w, problems.InternalError("Failed to retrieve device information"))
			return
		}
		userSessions = append(userSessions, data.Device{
			ID:         device.ID,
			UserID:     userID,
			FactoryID:  device.FactoryID,
			DeviceName: device.DeviceName,
			OsVersion:  device.OsVersion,
			CreatedAt:  device.CreatedAt,
			LastUsed:   device.LastUsed,
		})
	}

	err = Server.Queries.TerminateSessionsTransaction(r.Context(), userSessions, userID, httpkit.GenerateFingerprint(r), httpkit.GetClientIP(r))
	if err != nil {
		if errors.Is(err, data.ErrorDeviceDoesNotBelongToUser) {
			httpkit.RenderErr(w, problems.Forbidden("Device does not belong to user"))
			return
		}
		httpkit.RenderErr(w, problems.InternalError("Failed to terminate sessions"))
		return
	}

	httpkit.Render(w, http.StatusOK)
}

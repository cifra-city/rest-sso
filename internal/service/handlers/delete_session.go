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
	"github.com/cifra-city/rest-sso/resources"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
)

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDeleteSession(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	sessionID := req.Data.Attributes.DeviceId

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

	log.Infof("userID: %v", userID)

	_, err = Server.Queries.GetUserByID(r.Context(), userID)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve user information"))
		return
	}

	err = Server.Queries.DeleteDeviceByUserIdAndId(r.Context(), data.DeleteDeviceByUserIdAndIdParams{
		UserID: userID,
		ID:     uuid.MustParse(sessionID),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.NotFound("Device not found"))
			return
		}
		httpkit.RenderErr(w, problems.InternalError("Failed to delete device"))
		return
	}

	err = Server.Queries.InsertOperationHistory(r.Context(), data.InsertOperationHistoryParams{
		ID:            uuid.New(),
		UserID:        userID,
		DeviceData:    fingerprint,
		Operation:     data.OperationTypeDeleteSession,
		Success:       false,
		FailureReason: data.FailureReasonSuccess,
		IpAddress:     IP,
	})

	devices, err := Server.Queries.GetDevicesByUserID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.NotFound("Devices not found"))
			return
		}
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve devices"))
		return
	}
	var userSessions []resources.UserSessionDataAttributesDevicesInner
	for _, device := range devices {
		userSessions = append(userSessions, resources.UserSessionDataAttributesDevicesInner{
			Id:         device.ID.String(),
			FactoryId:  device.FactoryID,
			DeviceName: device.DeviceName.String,
			OsVersion:  device.OsVersion.String,
			LastUsed:   device.LastUsed,
		})
	}

	httpkit.Render(w, resources.UserSessions{
		Data: resources.UserSessionData{
			Type: "user_sessions",
			Attributes: resources.UserSessionDataAttributes{
				Devices: userSessions,
			},
		},
	})

}

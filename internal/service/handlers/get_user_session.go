package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/resources"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
)

func GetUserSessions(w http.ResponseWriter, r *http.Request) {
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

	devices, err := Server.Databaser.GetDevicesByUserID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.NotFound("No devices found"))
			return
		}
		log.Errorf("Failed to retrieve devices: %v", err)
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

package handlers

import (
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/resources"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetSessions(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Server](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration %s", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := Server.Logger

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	sessions, err := Server.Databaser.Sessions.GetSessions(r, userID)
	if err != nil {
		log.Errorf("Failed to retrieve user sessions: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	var userSessions []resources.UserSessionsDataAttributesDevicesInner
	for _, device := range sessions {
		userSessions = append(userSessions, resources.UserSessionsDataAttributesDevicesInner{
			Id:         device.ID.String(),
			Location:   "TODO",
			DeviceName: "TODO",
			Client:     "TODO",
			LastUsed:   device.LastUsed,
		})
	}

	httpkit.Render(w, resources.UserSessions{
		Data: resources.UserSessionsData{
			Type: resources.UserSessionType,
			Attributes: resources.UserSessionsDataAttributes{
				Devices: userSessions,
			},
		},
	})
}

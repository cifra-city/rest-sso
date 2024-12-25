package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
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

	sessionIdStr := req.Data.Attributes.SessionId

	sessionID, err := uuid.Parse(sessionIdStr)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("invalid format session id"))...)
		return
	}

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

	err = Server.Databaser.Sessions.Delete(r, sessionID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.NotFound("Device not found"))
			return
		}
		httpkit.RenderErr(w, problems.InternalError("Failed to delete device"))
		return
	}

	sessions, err := Server.Databaser.Sessions.GetSessions(r, userID)
	if err != nil {
		log.Errorf("Failed to retrieve user sessions: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	var userSessions []resources.UserSessionDataAttributesDevicesInner
	for _, device := range sessions {
		userSessions = append(userSessions, resources.UserSessionDataAttributesDevicesInner{
			Id:         device.ID.String(),
			DeviceName: "TODO",
			Client:     "TODO",

			LastUsed: device.LastUsed,
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

	httpkit.Render(w, resources.UserSessions{
		Data: resources.UserSessionData{
			Type: "user_sessions",
			Attributes: resources.UserSessionDataAttributes{
				Devices: userSessions,
			},
		},
	})

}

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
	"github.com/sirupsen/logrus"
)

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDeleteSession(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	sessionIdStr := req.Data.Attributes.SessionId

	sessionForDelete, err := uuid.Parse(sessionIdStr)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration %s", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := Server.Logger

	sessionID, ok := r.Context().Value(tokens.DeviceIDKey).(uuid.UUID)
	if !ok {
		log.Warn("SessionID not found in context")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if sessionID == sessionForDelete {
		log.Debugf("Session can't be current")
		httpkit.RenderErr(w, problems.BadRequest(errors.New("session can't be current"))...)
		return
	}

	err = Server.Databaser.Sessions.Delete(r, sessionForDelete, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.Errorf("Failed to delete device: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.TokenManager.Bin.Add(userID.String(), sessionForDelete.String())
	if err != nil {
		log.Errorf("Failed to add token to bin: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
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
}

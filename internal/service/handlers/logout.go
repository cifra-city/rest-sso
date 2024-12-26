package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	logrus.Debugf("user logged out")

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError("Failed to retrieve service configuration"))
		return
	}

	log := Server.Logger

	sessionID, ok := r.Context().Value(tokens.DeviceIDKey).(uuid.UUID)
	if !ok {
		log.Warn("SessionID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("Session not authenticated"))
		return
	}

	userID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	err = Server.Databaser.Sessions.Delete(r, sessionID, userID)

	httpkit.Render(w, http.StatusOK)
}

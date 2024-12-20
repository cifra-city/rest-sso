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
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func ResetPasswordComplete(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewResetPasswordComplete(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	username := req.Data.Attributes.Username
	firstPassword := req.Data.Attributes.FirstPassword
	secondPassword := req.Data.Attributes.SecondPassword

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	if email == nil && username == nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("email or username is required"))...)
		return
	}

	if firstPassword != secondPassword {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("passwords do not match"))...)
		return
	}

	var user data.UsersSecret

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	if username != nil {
		user, err = Server.Queries.GetUserByUsername(r.Context(), *username)
	} else {
		user, err = Server.Queries.GetUserByEmail(r.Context(), *email)
	}
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.Unauthorized())
			return
		}
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Mailman.CheckAccess(user.Email, string(RESET_PASSWORD), UserAgent, IP)
	if err != nil {
		log.Debugf("User %s has no access to reset password", user.Email)
		err = Server.Queries.InsertOperationHistory(r.Context(), data.InsertOperationHistoryParams{
			ID:            uuid.New(),
			UserID:        user.ID,
			DeviceData:    httpkit.GenerateFingerprint(r),
			Operation:     data.OperationTypeResetPassword,
			Success:       false,
			FailureReason: data.FailureReasonNoAccess,
			IpAddress:     IP,
		})
		httpkit.RenderErr(w, problems.Forbidden())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(firstPassword), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error hashing password: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Queries.ResetPasswordTransaction(r.Context(), &user, string(hashedPassword), httpkit.GenerateFingerprint(r), IP)
	if err != nil {
		log.Errorf("error make transaction reset pasword: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("user logged in: %s", user.Username)

	httpkit.Render(w, http.StatusAccepted)
}

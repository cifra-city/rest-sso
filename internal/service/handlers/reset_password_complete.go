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
	"github.com/cifra-city/rest-sso/internal/service/utils"
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

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	user, err := utils.GetUserExists(r.Context(), Server, username, email)
	if err != nil {
		if errors.Is(err, utils.ErrMultipleChoices) {
			log.Errorf("multiple choices error: %v", err)
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		if errors.Is(err, sql.ErrNoRows) {
			log.Errorf("user not found: %v", err)
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.Errorf("error getting user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Mailman.CheckAccess(user.Email, string(RESET_PASSWORD), UserAgent, IP)
	if err != nil {
		log.Debugf("User %s has no access to reset password", user.Email)
		err = Server.Databaser.InsertOperationHistory(r.Context(), dbcore.InsertOperationHistoryParams{
			ID:            uuid.New(),
			UserID:        user.ID,
			DeviceData:    httpkit.GenerateFingerprint(r),
			Operation:     dbcore.OperationTypeResetPassword,
			Success:       false,
			FailureReason: dbcore.FailureReasonNoAccess,
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

	err = Server.Databaser.ResetPasswordTransaction(r.Context(), &user, string(hashedPassword), httpkit.GenerateFingerprint(r), IP)
	if err != nil {
		log.Errorf("error make transaction reset pasword: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("user logged in: %s", user.Username)

	httpkit.Render(w, http.StatusAccepted)
}

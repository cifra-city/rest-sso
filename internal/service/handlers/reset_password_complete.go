package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/mailman"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/data/db/dbcore"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func ResetPasswordComplete(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Server](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration %s", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := Server.Logger

	req, err := requests.NewCredentials(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	password := req.Data.Attributes.Password

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	acc, err := Server.Databaser.Accounts.GetByEmail(r, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debugf("user not found for email: %v", email)
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.Errorf("error getting user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if !Server.Config.Email.Off { // for testing
		err = Server.Mailman.CheckAccess(acc.Email, string(RESET_PASSWORD), UserAgent, IP)
		if err != nil {
			_ = Server.Databaser.Operations.CreateFailure(r, acc.ID, dbcore.OperationTypeResetPassword, dbcore.FailureReasonNoAccess)
			if errors.Is(err, mailman.ErrNotFound) {
				log.Debugf("email haven`t request: %s", email)
				httpkit.RenderErr(w, problems.NotFound())
				return
			}
			if errors.Is(err, mailman.ErrAccessDenied) {
				log.Warnf("Metadata is invalid at try to reset password account: %s", acc.Email)
				httpkit.RenderErr(w, problems.Forbidden())
				return
			}
			log.Debugf("Access denied %s, %s %s", err, IP, UserAgent)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error hashing password: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Databaser.ResetPasswordTxn(r, acc.ID, string(hashedPassword))
	if err != nil {
		log.Errorf("error make transaction reset pasword: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, http.StatusAccepted)
}

package handlers

import (
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/mailman"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/cifra-city/rest-sso/internal/service/requests"
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
	password := req.Data.Attributes.Password

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	if email == nil && username == nil {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("email or username is required"))...)
		return
	}

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	acc, err := Server.Databaser.Accounts.Exists(r, username, email)
	if err != nil {
		log.Errorf("error getting user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	if acc == nil {
		httpkit.RenderErr(w, problems.NotFound())
		return
	}

	err = Server.Mailman.CheckAccess(acc.Email, string(RESET_PASSWORD), UserAgent, IP)
	if err != nil {
		_ = Server.Databaser.Operations.CreateFailure(r, acc.ID, dbcore.OperationTypeResetPassword, dbcore.FailureReasonNoAccess)
		if errors.Is(err, mailman.ErrNotFound) {
			log.Debugf("email haven`t request: %s", email)
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		if errors.Is(err, mailman.ErrAccessDenied) {
			log.Warnf("Metadata is invalid at try to reset password account: %s", acc.Username)
			httpkit.RenderErr(w, problems.Forbidden())
			return
		}
		log.Debugf("Access denied %s, %s %s", err, IP, UserAgent)
		httpkit.RenderErr(w, problems.InternalError())
		return
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

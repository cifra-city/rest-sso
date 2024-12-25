package handlers

import (
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/sectools"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationComplete(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistrationComplete(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	password := req.Data.Attributes.Password
	email := req.Data.Attributes.Email
	username := req.Data.Attributes.Username

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	if len(password) < 8 || !sectools.HasRequiredChars(password) {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("invalid password requirements"))...)
		return
	}

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpkit.RenderErr(w, problems.InternalError("database queries not found"))
		return
	}

	log := Server.Logger

	acc, err := Server.Databaser.Accounts.Exists(r, &username, &email)
	if err != nil {
		log.Errorf("error getting user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	if acc != nil {
		httpkit.RenderErr(w, problems.Conflict("user already exists"))
		return
	}

	err = Server.Mailman.CheckAccess(email, "registration", UserAgent, IP)
	if err != nil {
		if errors.Is(err, errors.New("not found")) {
			log.Warnf("email haven`t access: %s", email)
			httpkit.RenderErr(w, problems.NotFound("email haven`t access"))
			return
		}
		if errors.Is(err, errors.New("access denied")) {
			log.Warnf("failed to decrypt ConfidenceCode for email: %s", email)
			httpkit.RenderErr(w, problems.Forbidden("failed to decrypt ConfidenceCode"))
			return
		}
		log.Warnf("Access denied %s, %s %s", err, IP, UserAgent)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error hashing password: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	account, err := Server.Databaser.Accounts.Create(r, username, email, string(hashedPassword))
	if err != nil {
		logrus.Errorf("error inserting user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	logrus.Infof("user created: %v", account.Username)

	httpkit.Render(w, http.StatusCreated)
}

package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/mailman"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/sectools"
	"github.com/cifra-city/rest-sso/internal/service/events"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationComplete(w http.ResponseWriter, r *http.Request) {
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

	password := req.Data.Attributes.Password
	email := req.Data.Attributes.Email

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	if len(password) < 8 || !sectools.HasRequiredChars(password) || len(password) > 32 {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("invalid password requirements"))...)
		return
	}

	if !Server.Config.Email.Off { // for testing
		err = Server.Mailman.CheckAccess(email, string(REGISTRATION), UserAgent, IP)
		if err != nil {
			if errors.Is(err, mailman.ErrNotFound) {
				log.Debugf("email haven`t request: %s", email)
				httpkit.RenderErr(w, problems.NotFound())
				return
			}
			if errors.Is(err, mailman.ErrAccessDenied) {
				log.Warnf("Metadata is invalid at try to register account with : %s", email)
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
		log.Errorf("error hashing password: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	account, err := Server.Databaser.Accounts.Create(r, email, string(hashedPassword))
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"accounts_email_key\"" {
			httpkit.RenderErr(w, problems.Conflict())
			return
		}
		log.Errorf("error inserting account: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	event := events.AccountCreated{
		Event:     "AccountCreate",
		UserID:    account.ID.String(),
		Email:     email,
		Timestamp: time.Now().UTC(),
	}

	body, err := json.Marshal(event)
	if err != nil {
		http.Error(w, "Failed to serialize event", http.StatusInternalServerError)
		return
	}

	err = Server.Broker.Publish("sso.events", "account.create", "account.create", body)
	if err != nil {
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		return
	}

	log.Infof("account created: %v", account.Email)

	httpkit.Render(w, http.StatusCreated)
}

package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/sirupsen/logrus"
)

func RegistrationInitiate(w http.ResponseWriter, r *http.Request) {
	Server, err := cifractx.GetValue[*config.Server](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration %s", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := Server.Logger

	req, err := requests.NewEmail(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	_, err = Server.Databaser.Accounts.GetByEmail(r, email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("error getting user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	if err == nil {
		log.Debugf("user already exists for email: %v", email)
		httpkit.RenderErr(w, problems.Conflict())
		return
	}

	go func() {
		err = Server.Mailman.SendList(email, string(REGISTRATION), "email_list.html", UserAgent, IP, 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Debugf("Email sent successfully to: %s", email)
		}
	}()

	httpkit.Render(w, http.StatusAccepted)
}

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
	"github.com/sirupsen/logrus"
)

func ResetPasswordInitiate(w http.ResponseWriter, r *http.Request) {
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

	go func() {
		err = Server.Mailman.SendList(acc.Email, string(RESET_PASSWORD), "email_list.html", UserAgent, IP, 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Debugf("Email sent successfully to: %s", acc.Email)
		}
	}()

	httpkit.Render(w, http.StatusOK)
}

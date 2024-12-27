package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/sirupsen/logrus"
)

func ResetPasswordInitiate(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewResetPasswordInitiate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	username := req.Data.Attributes.Username

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
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

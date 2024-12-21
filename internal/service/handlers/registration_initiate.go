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
	"github.com/cifra-city/rest-sso/internal/service/utils"
	"github.com/sirupsen/logrus"
)

func RegistrationInitiate(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistrationInitiate(r)
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
		logrus.Errorf("error getting db queries: %v", err)
		httpkit.RenderErr(w, problems.InternalError("database queries not found"))
		return
	}

	log := Server.Logger

	_, err = utils.GetUserExists(r.Context(), Server, &username, &email)
	if !errors.Is(err, sql.ErrNoRows) {
		if err != nil {
			log.Errorf("error getting user: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
		httpkit.RenderErr(w, problems.Conflict("user already exists"))
		return
	}

	go func() {
		err = Server.Mailman.SendList(email, string(REGISTRATION), "email_list.html", UserAgent, IP, 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Infof("Email sent successfully to: %s", email)
		}
	}()

	httpkit.Render(w, http.StatusAccepted)
}

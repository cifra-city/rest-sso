package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/sirupsen/logrus"
)

func RegistrationInitiate(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistrationInitiate(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	em := req.Data.Attributes.Email
	username := req.Data.Attributes.Username

	IP := httpresp.GetClientIP(r)
	UserAgent := httpresp.GetUserAgent(r)

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpresp.RenderErr(w, problems.InternalError("database queries not found"))
		return
	}

	log := Server.Logger

	_, err = Server.Queries.GetUserByEmail(r.Context(), em)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("error getting user by email: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	if err == nil {
		httpresp.RenderErr(w, problems.Conflict("this email address already exists"))
		return
	}

	_, err = Server.Queries.GetUserByUsername(r.Context(), username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("error getting user by username: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	if err == nil {
		httpresp.RenderErr(w, problems.Conflict("this username already exists"))
		return
	}

	go func() {
		err = Server.Mailman.SendList(em, string(REGISTRATION), "email_list.html", UserAgent, IP, 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Infof("Email sent successfully to: %s", em)
		}
	}()

	httpresp.Render(w, http.StatusAccepted)
}

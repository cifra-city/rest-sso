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

func Registration(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRegistration(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	em := req.Data.Attributes.Email
	username := *req.Data.Attributes.Username

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		http.Error(w, "Database queries not found", http.StatusInternalServerError)
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

	err = Server.Mailman.SendList(username, "registration", em)
	if err != nil {
		log.Errorf("error sending email: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	httpresp.Render(w, http.StatusFound)
}

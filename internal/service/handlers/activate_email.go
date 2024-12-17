package handlers

import (
	"net/http"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/sirupsen/logrus"
)

func ActivateEmail(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewActivateEmail(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	code := req.Data.Attributes.Code

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		http.Error(w, "Database queries not found", http.StatusInternalServerError)
		return
	}

	log := Server.Logger

	if Server.Mailman.CheckCode(email, "registration", code) {
		Server.Mailman.AddAccessForUser(email, "registration")

		log.Debugf("code is correct add access for email: %s", email)
		httpresp.Render(w, http.StatusOK)
		return
	}

	log.Debugf("code is incorrect for email: %s", email)
	httpresp.Render(w, http.StatusForbidden)
}

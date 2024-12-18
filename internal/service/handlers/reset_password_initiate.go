package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/sirupsen/logrus"
)

func ResetPasswordInitiate(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewResetPasswordInitiate(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	username := req.Data.Attributes.Username

	var user data.UsersSecret

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	if username != nil {
		user, err = Server.Queries.GetUserByUsername(r.Context(), *username)
	} else {
		user, err = Server.Queries.GetUserByEmail(r.Context(), *email)
	}
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			httpresp.RenderErr(w, problems.NotFound())
			return
		}
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	go func() {
		err = Server.Mailman.SendList(user.Email, string(RESET_PASSWORD), "email_list.html", 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Infof("Email sent successfully to: %s", user.Email)
		}
	}()

	httpresp.Render(w, http.StatusOK)
}

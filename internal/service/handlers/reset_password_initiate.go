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
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	user, err := utils.GetUserExists(r.Context(), Server, username, email)
	if err != nil {
		if errors.Is(err, utils.ErrMultipleChoices) {
			log.Errorf("multiple choices error: %v", err)
			httpkit.RenderErr(w, problems.BadRequest(err)...)
			return
		}
		if errors.Is(err, sql.ErrNoRows) {
			log.Errorf("user not found: %v", err)
			httpkit.RenderErr(w, problems.NotFound())
			return
		}
		log.Errorf("error getting user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	go func() {
		err = Server.Mailman.SendList(user.Email, string(RESET_PASSWORD), "email_list.html", UserAgent, IP, 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Infof("Email sent successfully to: %s", user.Email)
		}
	}()

	httpkit.Render(w, http.StatusOK)
}

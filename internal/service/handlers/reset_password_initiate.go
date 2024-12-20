package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
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

	var user data.UsersSecret
	var user2 data.UsersSecret

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting server from context: %v", err)
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	if username != nil {
		user, err = Server.Queries.GetUserByUsername(r.Context(), *username)
	}
	if email != nil {
		user2, err = Server.Queries.GetUserByEmail(r.Context(), *email)
	}
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.Unauthorized())
			return
		}
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	if user2.ID != user.ID {
		httpkit.RenderErr(w, problems.BadRequest(errors.New("email and username do not match"))...)
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

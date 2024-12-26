package handlers

import (
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func LoginInitiate(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewLoginInitiate(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	username := req.Data.Attributes.Username
	password := req.Data.Attributes.Password

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpkit.RenderErr(w, problems.InternalError("database queries not found"))
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
		log.Debugf("user not found for email: %v", &email)
		httpkit.RenderErr(w, problems.NotFound())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.PassHash), []byte(password))
	if err != nil {
		err = Server.Databaser.Operations.CreateFailure(r, acc.ID, dbcore.OperationTypeLogin, dbcore.FailureReasonInvalidPassword)
		log.Debugf("Incorrect password for user: %s, error: %s", acc.Username, err)
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	go func() {
		err = Server.Mailman.SendList(acc.Email, string(LOGIN), "email_list.html", UserAgent, IP, 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Infof("Email sent successfully to: %s", email)
		}
	}()

	httpkit.Render(w, http.StatusAccepted)
}

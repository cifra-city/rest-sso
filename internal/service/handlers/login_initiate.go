package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/data/db/dbcore"
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
	password := req.Data.Attributes.Password

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

	Server, err := cifractx.GetValue[*config.Server](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	log := Server.Logger

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

	err = bcrypt.CompareHashAndPassword([]byte(acc.PassHash), []byte(password))
	if err != nil {
		err = Server.Databaser.Operations.CreateFailure(r, acc.ID, dbcore.OperationTypeLogin, dbcore.FailureReasonInvalidPassword)
		log.Debugf("Incorrect password for account: %s, error: %s", acc.Email, err)
		httpkit.RenderErr(w, problems.Conflict())
		return
	}

	go func() {
		err = Server.Mailman.SendList(acc.Email, string(LOGIN), "email_list.html", UserAgent, IP, 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Debugf("Email sent successfully to: %s", email)
		}
	}()

	httpkit.Render(w, http.StatusAccepted)
}

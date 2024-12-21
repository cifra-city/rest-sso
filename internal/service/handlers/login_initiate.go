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
	"github.com/google/uuid"
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
	fingerprint := httpkit.GenerateFingerprint(r)

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("error getting db queries: %v", err)
		httpkit.RenderErr(w, problems.InternalError("database queries not found"))
		return
	}

	log := Server.Logger

	var user data.UsersSecret
	var user2 data.UsersSecret

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

	err = bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password))
	if err != nil {
		err = Server.Queries.InsertOperationHistory(r.Context(), data.InsertOperationHistoryParams{
			ID:            uuid.New(),
			UserID:        user.ID,
			DeviceData:    fingerprint,
			Operation:     data.OperationTypeLogin,
			Success:       false,
			FailureReason: data.FailureReasonInvalidPassword,
			IpAddress:     IP,
		})

		log.Debugf("Incorrect password for user: %s, error: %s", user.Username, err)
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	go func() {
		err = Server.Mailman.SendList(user.Email, string(LOGIN), "email_list.html", UserAgent, IP, 300)
		if err != nil {
			log.Errorf("error sending email: %v", err)
		} else {
			log.Infof("Email sent successfully to: %s", email)
		}
	}()

	httpkit.Render(w, http.StatusAccepted)
}

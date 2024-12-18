package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/cifrajwt"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/cifra-city/rest-sso/pkg/sectools"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func ResetPasswordComplete(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewResetPasswordComplete(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	username := req.Data.Attributes.Username
	firstPassword := req.Data.Attributes.FirstPassword
	secondPassword := req.Data.Attributes.SecondPassword

	if email == nil && username == nil {
		httpresp.RenderErr(w, problems.BadRequest(errors.New("email or username is required"))...)
		return
	}

	if firstPassword != secondPassword {
		httpresp.RenderErr(w, problems.BadRequest(errors.New("passwords do not match"))...)
		return
	}

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
			httpresp.RenderErr(w, problems.Unauthorized())
			return
		}
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	if !Server.Mailman.CheckAndDeleteAccessForUser(user.Email, string(RESET_PASSWORD)) {
		log.Debugf("User %s has no access to reset password", user.Email)
		httpresp.RenderErr(w, problems.Forbidden())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(firstPassword), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("error hashing password: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	tokenRefresh, err := cifrajwt.GenerateJWT(user.ID, string(user.Role), int(user.TokenVersion+1), Server.Config.JWT.RefreshToken.TokenLifetime, Server.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		log.Errorf("error generating token access jwt: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	expiresAt := time.Now().UTC().Add(Server.Config.JWT.RefreshToken.TokenLifetime)

	encryptedToken, err := sectools.EncryptToken(tokenRefresh, Server.Config.JWT.RefreshToken.EncryptionKey)
	if err != nil {
		log.Errorf("Failed to encrypt refresh token: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Queries.ResetPassword(r.Context(), &user, encryptedToken, expiresAt, string(hashedPassword))
	if err != nil {
		log.Errorf("error make transaction reset pasword: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("user logged in: %s", user.Username)

	httpresp.Render(w, http.StatusAccepted)
}

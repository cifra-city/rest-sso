package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/mailman"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/sectools"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/resources"
	"github.com/cifra-city/tokens"
	"github.com/sirupsen/logrus"
)

func LoginComplete(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewLoginComplete(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	username := req.Data.Attributes.Username
	factoryId := req.Data.Attributes.FactoryId
	deviceName := req.Data.Attributes.DeviceName
	osVersion := req.Data.Attributes.OsVersion

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

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

	err = Server.Mailman.CheckAccess(user.Email, string(LOGIN), UserAgent, IP)
	if err != nil {
		if errors.Is(err, mailman.ErrNotFound) {
			log.Warnf("email haven`t access: %s", user.Email)
			httpkit.RenderErr(w, problems.NotFound("email haven`t access"))
			return
		}
		if errors.Is(err, mailman.ErrAccessDenied) {
			log.Warnf("failed to decrypt ConfidenceCode for email: %s", user.Email)
			httpkit.RenderErr(w, problems.Forbidden("failed to decrypt ConfidenceCode"))
			return
		}
		log.Warnf("Access denied %s, %s %s", err, IP, UserAgent)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	tokenAccess, err := tokens.GenerateJWT(user.ID, string(user.Role), int(user.TokenVersion), Server.Config.JWT.AccessToken.TokenLifetime, Server.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		log.Errorf("error generating token access jwt: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	tokenRefresh, err := tokens.GenerateJWT(user.ID, string(user.Role), int(user.TokenVersion), Server.Config.JWT.RefreshToken.TokenLifetime, Server.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		log.Errorf("error generating token access jwt: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	expiresAt := time.Now().UTC().Add(Server.Config.JWT.RefreshToken.TokenLifetime)

	encryptedToken, err := sectools.EncryptToken(tokenRefresh, Server.Config.JWT.RefreshToken.EncryptionKey)
	if err != nil {
		log.Errorf("Failed to encrypt refresh token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Queries.UpdateRefreshTokenTransaction(r.Context(), &user, factoryId, deviceName, osVersion, encryptedToken, expiresAt, IP)
	if err != nil {
		log.Errorf("error updating last used and refresh token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("user logged in: %s", user.Username)

	logrus.Infof("user created: %v", user.Username)
	httpkit.Render(w, resources.LoginCompleteResp{
		Data: resources.LoginCompleteRespData{
			Type: "login",
			Attributes: resources.LoginCompleteRespDataAttributes{
				AccessToken:  tokenAccess,
				RefreshToken: tokenRefresh,
				ExpiresIn:    int32(Server.Config.JWT.AccessToken.TokenLifetime.Seconds()),
			},
		},
	})
}

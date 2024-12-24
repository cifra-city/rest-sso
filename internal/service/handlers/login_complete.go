package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/mailman"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/sectools"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/internal/service/utils"
	"github.com/cifra-city/rest-sso/resources"
	"github.com/google/uuid"
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

	deviceID := uuid.New()

	tokenAccess, tokenRefresh, expiresAt, err := utils.GenerateTokens(*Server, user, deviceID)
	if err != nil {
		log.Errorf("error generating tokens: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	encryptToken, err := sectools.EncryptToken(tokenRefresh, Server.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		log.Errorf("error encrypting token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Databaser.LoginTransaction(r.Context(), user.ID, deviceID, encryptToken, expiresAt,
		factoryId, deviceName, osVersion, IP, httpkit.GenerateFingerprint(r))
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

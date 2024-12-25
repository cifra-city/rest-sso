package handlers

import (
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
	deviceName := req.Data.Attributes.DeviceName

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
		log.Debugf("user not found: %v", err)
		httpkit.RenderErr(w, problems.NotFound())
		return
	}

	err = Server.Mailman.CheckAccess(acc.Email, string(LOGIN), UserAgent, IP)
	if err != nil {
		if errors.Is(err, mailman.ErrNotFound) {
			log.Warnf("email haven`t access: %s", acc.Email)
			httpkit.RenderErr(w, problems.NotFound("email haven`t access"))
			return
		}
		if errors.Is(err, mailman.ErrAccessDenied) {
			log.Warnf("failed to decrypt ConfidenceCode for email: %s", acc.Email)
			httpkit.RenderErr(w, problems.Forbidden("failed to decrypt ConfidenceCode"))
			return
		}
		log.Warnf("Access denied %s, %s %s", err, IP, UserAgent)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	deviceID := uuid.New()

	tokenAccess, tokenRefresh, _, err := utils.GenerateTokens(*Server, *acc, deviceID)
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

	_, err = Server.Databaser.LoginTxn(r, acc.ID, encryptToken, deviceName)
	if err != nil {
		log.Errorf("error updating last used and refresh token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("user logged in: %s", acc.Username)

	httpkit.Render(w, resources.LoginCompleteResp{
		Data: resources.LoginCompleteRespData{
			Type: "login",
			Attributes: resources.LoginCompleteRespDataAttributes{
				AccessToken:  tokenAccess,
				RefreshToken: tokenRefresh,
			},
		},
	})
}

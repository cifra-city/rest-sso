package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/comtools/cifractx"
	"github.com/cifra-city/comtools/httpkit"
	"github.com/cifra-city/comtools/httpkit/problems"
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
	Server, err := cifractx.GetValue[*config.Server](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration %s", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := Server.Logger

	req, err := requests.NewLogin(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	email := req.Data.Attributes.Email
	deviceName := req.Data.Attributes.DeviceName

	IP := httpkit.GetClientIP(r)
	UserAgent := httpkit.GetUserAgent(r)

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

	if !Server.Config.Email.Off { // for testing
		err = Server.Mailman.CheckAccess(acc.Email, string(LOGIN), UserAgent, IP)
		if err != nil {
			if errors.Is(err, mailman.ErrNotFound) {
				log.Infof("email haven`t access: %s", acc.Email)
				httpkit.RenderErr(w, problems.NotFound())
				return
			}
			if errors.Is(err, mailman.ErrAccessDenied) {
				log.Warnf("failed to decrypt ConfidenceCode for email: %s", acc.Email)
				httpkit.RenderErr(w, problems.Forbidden())
				return
			}
			log.Infof("Access denied %s, %s %s", err, IP, UserAgent)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
	}

	deviceID := uuid.New()

	tokenAccess, tokenRefresh, err := utils.GenerateTokens(*Server, acc, deviceID)
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

	_, err = Server.Databaser.LoginTxn(r, acc.ID, deviceName, encryptToken, deviceID)
	if err != nil {
		log.Errorf("Error transaction login: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	Server.Mailman.DeleteAccess(acc.Email, string(LOGIN))

	httpkit.Render(w, resources.TokensPair{
		Data: resources.TokensPairData{
			Type: resources.TokensPairType,
			Attributes: resources.TokensPairDataAttributes{
				AccessToken:  tokenAccess,
				RefreshToken: tokenRefresh,
			},
		},
	})
}

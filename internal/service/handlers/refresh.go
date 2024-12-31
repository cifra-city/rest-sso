package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/httpkit/problems"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/data/db/dbcore"
	"github.com/cifra-city/rest-sso/internal/sectools"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/resources"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRefresh(r)
	if err != nil {
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Server, err := cifractx.GetValue[*config.Service](r.Context(), config.SERVICE)
	if err != nil {
		logrus.Errorf("Failed to get service from context: %v", err)
		http.Error(w, "Service configuration not found", http.StatusInternalServerError)
		return
	}
	log := Server.Logger

	refreshToken := req.Data.Attributes.RefreshToken

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Debugf("Missing Authorization header")
		httpkit.RenderErr(w, problems.Unauthorized("Missing Authorization header"))
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		log.Debugf("Invalid Authorization header format")
		httpkit.RenderErr(w, problems.Unauthorized("Invalid Authorization header format"))
		return
	}
	tokenString := parts[1]

	userData, err := Server.TokenManager.VerifyJWTAndExtractClaims(tokenString, Server.Config.JWT.AccessToken.SecretKey)
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		log.Warnf("Token validation failed: %v", err)
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	userID := userData.ID
	sessionID := userData.DevID

	user, err := Server.Databaser.Accounts.GetById(r, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.Unauthorized())
			return
		}
		log.Errorf("Failed to get user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	session, err := Server.Databaser.Sessions.GetByID(r, sessionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.Unauthorized("session not found"))
			return
		}
		log.Errorf("Failed to get session: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if session.UserID != userID {
		log.Warn("Device does not belong to user")
		err = Server.Databaser.Operations.CreateFailure(r, userID, dbcore.OperationTypeRefreshToken, dbcore.FailureReasonInvalidDeviceID)
		if err != nil {
			log.Errorf("Failed to create operation history: %v", err)
			httpkit.RenderErr(w, problems.InternalError())
			return
		}
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}

	decryptedToken, err := sectools.DecryptToken(session.Token, Server.Config.JWT.RefreshToken.EncryptionKey)
	if err != nil {
		log.Errorf("Failed to decrypt refresh token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if decryptedToken != refreshToken {
		Server.Logger.Warn("Provided refresh token does not match the stored token")
		httpkit.RenderErr(w, problems.Conflict())
		return
	}

	tokenAccess, err := Server.TokenManager.GenerateJWT(user.ID, sessionID, string(user.Role), Server.Config.JWT.AccessToken.TokenLifetime, Server.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		Server.Logger.Errorf("Error generating access token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	tokenRefresh, err := Server.TokenManager.GenerateJWT(user.ID, sessionID, string(user.Role), Server.Config.JWT.RefreshToken.TokenLifetime, Server.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		Server.Logger.Errorf("Error generating refresh token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	encryptedToken, err := sectools.EncryptToken(tokenRefresh, Server.Config.JWT.RefreshToken.EncryptionKey)
	if err != nil {
		log.Errorf("Failed to encrypt refresh token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	err = Server.Databaser.UpdateRefreshTokenTrx(r, userID, sessionID, encryptedToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.Unauthorized())
			return
		}
		log.Errorf("Error updating last used and refresh token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, resources.RefreshResp{
		Data: resources.RefreshRespData{
			Type: string(dbcore.OperationTypeRefreshToken),
			Attributes: resources.RefreshRespDataAttributes{
				AccessToken:  tokenAccess,
				RefreshToken: tokenRefresh,
			},
		},
	})
}

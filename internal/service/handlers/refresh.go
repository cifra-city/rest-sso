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
	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/cifra-city/rest-sso/internal/sectools"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/resources"
	"github.com/cifra-city/tokens"
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
		log.Warn("Missing Authorization header")
		httpkit.RenderErr(w, problems.Unauthorized("Missing Authorization header"))
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		log.Warn("Invalid Authorization header format")
		httpkit.RenderErr(w, problems.Unauthorized("Invalid Authorization header format"))
		return
	}
	tokenString := parts[1]

	log.Debugf("Token received: %s", tokenString)

	userID, sessionID, tokenVersion, _, err := tokens.VerifyJWTAndExtractClaims(r.Context(), tokenString, Server.Config.JWT.AccessToken.SecretKey, log)
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		log.Warnf("Token validation failed: %v", err)
		httpkit.RenderErr(w, problems.Unauthorized("Token validation failed"))
		return
	}

	user, err := Server.Databaser.Accounts.GetById(r, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.Unauthorized("User not found"))
			return
		}
		log.Errorf("Failed to get user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	if int(user.TokenVersion) != tokenVersion {
		log.Warn("Token version mismatch")
		httpkit.RenderErr(w, problems.Unauthorized("Token version mismatch"))
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
		httpkit.RenderErr(w, problems.Unauthorized("Device does not belong to user"))
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
		httpkit.RenderErr(w, problems.Unauthorized("Invalid refresh token"))
		return
	}

	tokenAccess, err := tokens.GenerateJWT(user.ID, sessionID, string(user.Role), int(user.TokenVersion), Server.Config.JWT.AccessToken.TokenLifetime, Server.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		Server.Logger.Errorf("Error generating access token: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	tokenRefresh, err := tokens.GenerateJWT(user.ID, sessionID, string(user.Role), int(user.TokenVersion), Server.Config.JWT.RefreshToken.TokenLifetime, Server.Config.JWT.RefreshToken.SecretKey)
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
		log.Errorf("Error updating last used and refresh token: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			httpkit.RenderErr(w, problems.Unauthorized())
			return
		}
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("User logged in: %s", user.Username)

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

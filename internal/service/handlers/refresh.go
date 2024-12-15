package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/rest-sso/internal/service/requests"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/cifrajwt"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/cifra-city/rest-sso/pkg/httpresp/problems"
	"github.com/cifra-city/rest-sso/resources"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRefresh(r)
	if err != nil {
		httpresp.RenderErr(w, problems.BadRequest(err)...)
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
	deviceIdStr := req.Data.Attributes.DeviceId
	factoryId := req.Data.Attributes.FactoryId
	deviceName := req.Data.Attributes.DeviceName
	osVersion := req.Data.Attributes.OsVersion
	ipAddress := req.Data.Attributes.IpAddress

	//JWT Middleware

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Warn("Missing Authorization header")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		log.Warn("Invalid Authorization header format")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	tokenString := parts[1]

	log.Debugf("Token received: %s", tokenString)

	claims := &cifrajwt.CustomClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Server.Config.JWT.AccessToken.SecretKey), nil
	})

	if err != nil {
		log.Warnf("Invalid token: %v", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		log.Errorf("Invalid user ID in token claims: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Infof("Claims Subject (UserID): %s", userID)

	//END JWT Middleware

	user, err := Server.Queries.GetUserByID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpresp.RenderErr(w, problems.Unauthorized("User not found"))
			return
		}
		log.Errorf("Failed to get user: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	deviceId, err := uuid.Parse(deviceIdStr)
	if err != nil {
		log.Errorf("Invalid device ID format: %v", err)
		httpresp.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	log.Infof("Device ID: %s", deviceId)
	device, err := Server.Queries.GetDeviceByID(r.Context(), deviceId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			httpresp.RenderErr(w, problems.Unauthorized("device not found"))
			return
		}
		log.Errorf("Failed to get device: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	if device.UserID != userID {
		log.Warn("Device does not belong to user")
		_ = Server.Queries.InsertLoginHistory(r.Context(), data.InsertLoginHistoryParams{
			ID:        uuid.New(),
			UserID:    userID,
			DeviceID:  deviceId,
			IpAddress: ipAddress,
			LoginTime: time.Now().UTC(),
			Success:   false,
			FailureReason: data.NullFailureReason{
				FailureReason: data.FailureReasonInvalidDeviceID,
				Valid:         true,
			},
		})
		httpresp.RenderErr(w, problems.Unauthorized("Device does not belong to user"))
		return
	}

	if device.FactoryID != factoryId {
		log.Warn("Factory ID does not match")
		_ = Server.Queries.InsertLoginHistory(r.Context(), data.InsertLoginHistoryParams{
			ID:        uuid.New(),
			UserID:    userID,
			DeviceID:  deviceId,
			IpAddress: ipAddress,
			LoginTime: time.Now().UTC(),
			Success:   false,
			FailureReason: data.NullFailureReason{
				FailureReason: data.FailureReasonInvalidDeviceFactoryID,
				Valid:         true,
			},
		})
		httpresp.RenderErr(w, problems.Unauthorized("Factory ID does not match"))
		return
	}

	dbToken, err := Server.Queries.GetTokenByUserIdAndDeviceId(r.Context(), data.GetTokenByUserIdAndDeviceIdParams{
		UserID:   userID,
		DeviceID: deviceId,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_ = Server.Queries.InsertLoginHistory(r.Context(), data.InsertLoginHistoryParams{
				ID:        uuid.New(),
				UserID:    userID,
				DeviceID:  deviceId,
				IpAddress: ipAddress,
				LoginTime: time.Now().UTC(),
				Success:   false,
				FailureReason: data.NullFailureReason{
					FailureReason: data.FailureReasonInvalidRefreshToken,
				},
			})
			httpresp.RenderErr(w, problems.Unauthorized("Token not found"))
			return
		}
		Server.Logger.Errorf("Failed to fetch token from database: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	if dbToken.Token != refreshToken {
		Server.Logger.Warn("Provided refresh token does not match the stored token")
		httpresp.RenderErr(w, problems.Unauthorized("Invalid refresh token"))
		return
	}

	tokenAccess, err := cifrajwt.GenerateJWT(userID, Server.Config.JWT.AccessToken.TokenLifetime, Server.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		Server.Logger.Errorf("Error generating access token: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	tokenRefresh, err := cifrajwt.GenerateJWT(userID, Server.Config.JWT.RefreshToken.TokenLifetime, Server.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		Server.Logger.Errorf("Error generating refresh token: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	expiresAt := time.Now().UTC().Add(Server.Config.JWT.RefreshToken.TokenLifetime)

	err = Server.Queries.UpdateRefreshTokenTransaction(r.Context(), &user, device.FactoryID, deviceName, osVersion, tokenRefresh, expiresAt, ipAddress)
	if err != nil {
		log.Errorf("Error updating last used and refresh token: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("User logged in: %s", user.Username)

	// Отправляем новые токены клиенту
	httpresp.Render(w, resources.RefreshResp{
		Data: resources.RefreshRespData{
			Type: "refresh",
			Attributes: resources.RefreshRespDataAttributes{
				AccessToken:  tokenAccess,
				RefreshToken: tokenRefresh,
				ExpiresIn:    int32(Server.Config.JWT.AccessToken.TokenLifetime.Seconds()),
			},
		},
	})
}

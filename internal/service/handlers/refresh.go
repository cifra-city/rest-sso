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
	"github.com/cifra-city/rest-sso/pkg/sectools"
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

	IP := httpresp.GetClientIP(r)

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Warn("Missing Authorization header")
		httpresp.RenderErr(w, problems.Unauthorized("Missing Authorization header"))
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		log.Warn("Invalid Authorization header format")
		httpresp.RenderErr(w, problems.Unauthorized("Invalid Authorization header format"))
		return
	}
	tokenString := parts[1]

	log.Debugf("Token received: %s", tokenString)

	claims := &cifrajwt.CustomClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Server.Config.JWT.AccessToken.SecretKey), nil
	})

	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		log.Warnf("Invalid token: %v", err)
		httpresp.RenderErr(w, problems.Unauthorized("Invalid token"))
		return
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		log.Errorf("Invalid user ID in token claims: %v", err)
		httpresp.RenderErr(w, problems.Unauthorized("Invalid token"))
		return
	}

	log.Infof("Claims Subject (UserID): %s", userID)

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
		_ = Server.Queries.InsertOperationHistory(r.Context(), data.InsertOperationHistoryParams{
			ID:            uuid.New(),
			UserID:        userID,
			DeviceData:    httpresp.GenerateFingerprint(r),
			Operation:     data.OperationTypeRefreshToken,
			Success:       false,
			FailureReason: data.FailureReasonInvalidDeviceID,
			IpAddress:     IP,
		})
		httpresp.RenderErr(w, problems.Unauthorized("Device does not belong to user"))
		return
	}

	if device.FactoryID != factoryId {
		log.Warn("Factory ID does not match")
		_ = Server.Queries.InsertOperationHistory(r.Context(), data.InsertOperationHistoryParams{
			ID:            uuid.New(),
			UserID:        userID,
			DeviceData:    httpresp.GenerateFingerprint(r),
			Operation:     data.OperationTypeRefreshToken,
			Success:       false,
			FailureReason: data.FailureReasonInvalidDeviceFactoryID,
			IpAddress:     IP,
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
			_ = Server.Queries.InsertOperationHistory(r.Context(), data.InsertOperationHistoryParams{
				ID:            uuid.New(),
				UserID:        userID,
				DeviceData:    httpresp.GenerateFingerprint(r),
				Operation:     data.OperationTypeRefreshToken,
				Success:       false,
				FailureReason: data.FailureReasonInvalidRefreshToken,
				IpAddress:     IP,
			})
			httpresp.RenderErr(w, problems.Unauthorized("Token not found"))
			return
		}
		Server.Logger.Errorf("Failed to fetch token from database: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	decryptedToken, err := sectools.DecryptToken(dbToken.Token, Server.Config.JWT.RefreshToken.EncryptionKey)
	if err != nil {
		log.Errorf("Failed to decrypt refresh token: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	if decryptedToken != refreshToken {
		Server.Logger.Warn("Provided refresh token does not match the stored token")
		httpresp.RenderErr(w, problems.Unauthorized("Invalid refresh token"))
		return
	}

	tokenAccess, err := cifrajwt.GenerateJWT(user.ID, string(user.Role), int(user.TokenVersion), Server.Config.JWT.AccessToken.TokenLifetime, Server.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		Server.Logger.Errorf("Error generating access token: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}

	tokenRefresh, err := cifrajwt.GenerateJWT(user.ID, string(user.Role), int(user.TokenVersion), Server.Config.JWT.RefreshToken.TokenLifetime, Server.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		Server.Logger.Errorf("Error generating refresh token: %v", err)
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

	err = Server.Queries.UpdateRefreshTokenTransaction(r.Context(), &user, device.FactoryID, deviceName, osVersion, encryptedToken, expiresAt, IP)
	if err != nil {
		log.Errorf("Error updating last used and refresh token: %v", err)
		httpresp.RenderErr(w, problems.InternalError())
		return
	}
	log.Infof("User logged in: %s", user.Username)

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

package utils

import (
	"time"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
	"github.com/cifra-city/tokens"
	"github.com/google/uuid"
)

func GenerateTokens(service config.Service, user data.UsersSecret, deviceID uuid.UUID) (tokenAccess string, tokenRefresh string, expiresAt time.Time, err error) {
	tokenAccess, err = tokens.GenerateJWT(user.ID, deviceID, string(user.Role), int(user.TokenVersion), service.Config.JWT.AccessToken.TokenLifetime, service.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		return "", "", time.Time{}, err
	}

	tokenRefresh, err = tokens.GenerateJWT(user.ID, deviceID, string(user.Role), int(user.TokenVersion), service.Config.JWT.RefreshToken.TokenLifetime, service.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		return "", "", time.Time{}, err
	}

	expiresAt = time.Now().UTC().Add(service.Config.JWT.RefreshToken.TokenLifetime)

	return tokenAccess, tokenRefresh, expiresAt, nil
}

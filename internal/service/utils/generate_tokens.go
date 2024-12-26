package utils

import (
	"time"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/google/uuid"
)

func GenerateTokens(service config.Service, account dbcore.Account, deviceID uuid.UUID) (tokenAccess string, tokenRefresh string, expiresAt time.Time, err error) {
	tokenAccess, err = service.TokenManager.GenerateJWT(account.ID, deviceID, string(account.Role), int(account.TokenVersion), service.Config.JWT.AccessToken.TokenLifetime, service.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		return "", "", time.Time{}, err
	}

	tokenRefresh, err = service.TokenManager.GenerateJWT(account.ID, deviceID, string(account.Role), int(account.TokenVersion), service.Config.JWT.RefreshToken.TokenLifetime, service.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		return "", "", time.Time{}, err
	}

	expiresAt = time.Now().UTC().Add(service.Config.JWT.RefreshToken.TokenLifetime)

	return tokenAccess, tokenRefresh, expiresAt, nil
}

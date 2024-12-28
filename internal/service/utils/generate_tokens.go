package utils

import (
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/google/uuid"
)

func GenerateTokens(service config.Service, account dbcore.Account, deviceID uuid.UUID) (tokenAccess string, tokenRefresh string, err error) {
	tokenAccess, err = service.TokenManager.GenerateJWT(account.ID, deviceID, string(account.Role), service.Config.JWT.AccessToken.TokenLifetime, service.Config.JWT.AccessToken.SecretKey)
	if err != nil {
		return "", "", err
	}

	tokenRefresh, err = service.TokenManager.GenerateJWT(account.ID, deviceID, string(account.Role), service.Config.JWT.RefreshToken.TokenLifetime, service.Config.JWT.RefreshToken.SecretKey)
	if err != nil {
		return "", "", err
	}

	return tokenAccess, tokenRefresh, nil
}

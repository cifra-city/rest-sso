package utils

import (
	"context"
	"errors"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/db/data"
)

var (
	ErrMultipleChoices = errors.New("multiple choices error")
)

func GetUserExists(ctx context.Context, server *config.Service, username *string, email *string) (user dbcore.UsersSecret, err error) {
	var user1 dbcore.UsersSecret
	var user2 dbcore.UsersSecret

	if username != nil {
		user1, err = server.Databaser.GetUserByUsername(ctx, *username)
	}
	if email != nil {
		user2, err = server.Databaser.GetUserByEmail(ctx, *email)
	}
	if err != nil {
		server.Logger.Errorf("Failed to get user: %v", err)
		return dbcore.UsersSecret{}, err
	}

	if user2.ID != user1.ID {
		return dbcore.UsersSecret{}, ErrMultipleChoices
	}

	return user1, nil
}

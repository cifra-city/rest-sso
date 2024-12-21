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

func GetUserExists(ctx context.Context, server *config.Service, username *string, email *string) (user data.UsersSecret, err error) {
	var user1 data.UsersSecret
	var user2 data.UsersSecret

	if username != nil {
		user1, err = server.Queries.GetUserByUsername(ctx, *username)
	}
	if email != nil {
		user2, err = server.Queries.GetUserByEmail(ctx, *email)
	}
	if err != nil {
		server.Logger.Errorf("Failed to get user: %v", err)
		return data.UsersSecret{}, err
	}

	if user2.ID != user1.ID {
		return data.UsersSecret{}, ErrMultipleChoices
	}

	return user1, nil
}

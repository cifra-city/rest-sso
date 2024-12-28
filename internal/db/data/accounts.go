package data

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
	"github.com/google/uuid"
)

type Accounts interface {
	Create(r *http.Request, username string, email string, passHash string) (dbcore.Account, error)

	Exists(r *http.Request, username *string, email *string) (*dbcore.Account, error)

	GetById(r *http.Request, id uuid.UUID) (dbcore.Account, error)
	GetByLogin(r *http.Request, email string, username string) (dbcore.Account, error)
	GetByEmail(r *http.Request, email string) (dbcore.Account, error)
	GetByUsername(r *http.Request, email string) (dbcore.Account, error)

	UpdatePassword(r *http.Request, id uuid.UUID, passHash string) (dbcore.Account, error)
	UpdateUsername(r *http.Request, id uuid.UUID, username string) (dbcore.Account, error)
}

type accounts struct {
	queries *dbcore.Queries
}

func NewAccount(queries *dbcore.Queries) Accounts {
	return &accounts{queries: queries}
}

func (a *accounts) Create(r *http.Request, username string, email string, passHash string) (dbcore.Account, error) {
	return a.queries.CreateAccount(r.Context(), dbcore.CreateAccountParams{
		Username: username,
		Email:    email,
		PassHash: passHash,
	})
}

func (a *accounts) Exists(r *http.Request, username *string, email *string) (*dbcore.Account, error) {
	var userByUsername *dbcore.Account
	var userByEmail *dbcore.Account

	if username != nil && *username != "" {
		result, err := a.queries.GetAccountByUsername(r.Context(), *username)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		if err == nil {
			userByUsername = &result
		}
	}

	if email != nil && *email != "" {
		result, err := a.queries.GetAccountByEmail(r.Context(), *email)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		if err == nil {
			userByEmail = &result
		}
	}

	if userByUsername == nil && userByEmail == nil {
		return nil, nil
	}

	if userByUsername != nil && userByEmail != nil {
		if userByUsername.ID != userByEmail.ID {
			return nil, errors.New("conflicting accounts found for username and email")
		}
	}

	if userByUsername != nil {
		return userByUsername, nil
	}
	return userByEmail, nil
}

func (a *accounts) GetById(r *http.Request, id uuid.UUID) (dbcore.Account, error) {
	return a.queries.GetAccountByID(r.Context(), id)
}

func (a *accounts) GetByLogin(r *http.Request, email string, username string) (dbcore.Account, error) {
	return a.queries.GetAccountByLogin(r.Context(), dbcore.GetAccountByLoginParams{
		Username: username,
		Email:    email,
	})
}

func (a *accounts) GetByEmail(r *http.Request, email string) (dbcore.Account, error) {
	return a.queries.GetAccountByEmail(r.Context(), email)
}

func (a *accounts) GetByUsername(r *http.Request, email string) (dbcore.Account, error) {
	return a.queries.GetAccountByUsername(r.Context(), email)
}

func (a *accounts) UpdatePassword(r *http.Request, id uuid.UUID, passHash string) (dbcore.Account, error) {
	return a.queries.UpdatePassword(r.Context(), dbcore.UpdatePasswordParams{
		ID:       id,
		PassHash: passHash,
	})
}

func (a *accounts) UpdateUsername(r *http.Request, id uuid.UUID, username string) (dbcore.Account, error) {
	return a.queries.UpdateUsername(r.Context(), dbcore.UpdateUsernameParams{
		ID:       id,
		Username: username,
	})
}

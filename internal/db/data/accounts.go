package data

import (
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
	UpdateTokenVersion(r *http.Request, id uuid.UUID) (dbcore.Account, error)
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
	var acc1 dbcore.Account
	var acc2 dbcore.Account
	var err error

	if username != nil {
		acc1, err = a.queries.GetAccountByUsername(r.Context(), *username)
	}
	if email != nil {
		acc2, err = a.queries.GetAccountByEmail(r.Context(), *email)
	}
	if err != nil {
		return nil, err
	}

	if acc1.ID != acc2.ID {
		return nil, nil
	}

	return &acc1, nil
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

func (a *accounts) UpdateTokenVersion(r *http.Request, id uuid.UUID) (dbcore.Account, error) {
	return a.queries.UpdateTokenVersion(r.Context(), id)
}

func (a *accounts) UpdateUsername(r *http.Request, id uuid.UUID, username string) (dbcore.Account, error) {
	return a.queries.UpdateUsername(r.Context(), dbcore.UpdateUsernameParams{
		ID:       id,
		Username: username,
	})
}

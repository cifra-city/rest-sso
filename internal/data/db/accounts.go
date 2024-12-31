package db

import (
	"net/http"

	"github.com/cifra-city/rest-sso/internal/data/db/dbcore"
	"github.com/google/uuid"
)

type Accounts interface {
	Create(r *http.Request, email string, passHash string) (dbcore.Account, error)

	GetById(r *http.Request, id uuid.UUID) (dbcore.Account, error)
	GetByEmail(r *http.Request, email string) (dbcore.Account, error)

	UpdatePassword(r *http.Request, id uuid.UUID, passHash string) (dbcore.Account, error)
}

type accounts struct {
	queries *dbcore.Queries
}

func NewAccount(queries *dbcore.Queries) Accounts {
	return &accounts{queries: queries}
}

func (a *accounts) Create(r *http.Request, email string, passHash string) (dbcore.Account, error) {
	return a.queries.CreateAccount(r.Context(), dbcore.CreateAccountParams{
		Email:    email,
		PassHash: passHash,
	})
}

func (a *accounts) GetById(r *http.Request, id uuid.UUID) (dbcore.Account, error) {
	return a.queries.GetAccountByID(r.Context(), id)
}

func (a *accounts) GetByEmail(r *http.Request, email string) (dbcore.Account, error) {
	return a.queries.GetAccountByEmail(r.Context(), email)
}

func (a *accounts) UpdatePassword(r *http.Request, id uuid.UUID, passHash string) (dbcore.Account, error) {
	return a.queries.UpdatePassword(r.Context(), dbcore.UpdatePasswordParams{
		ID:       id,
		PassHash: passHash,
	})
}

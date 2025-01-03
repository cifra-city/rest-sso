package db

import (
	"net/http"

	"github.com/cifra-city/rest-sso/internal/data/db/sqlcore"
	"github.com/google/uuid"
)

type Accounts interface {
	Create(r *http.Request, email string, passHash string) (sqlcore.Account, error)

	GetById(r *http.Request, id uuid.UUID) (sqlcore.Account, error)
	GetByEmail(r *http.Request, email string) (sqlcore.Account, error)

	UpdatePassword(r *http.Request, id uuid.UUID, passHash string) (sqlcore.Account, error)
}

type accounts struct {
	queries *sqlcore.Queries
}

func NewAccount(queries *sqlcore.Queries) Accounts {
	return &accounts{queries: queries}
}

func (a *accounts) Create(r *http.Request, email string, passHash string) (sqlcore.Account, error) {
	return a.queries.CreateAccount(r.Context(), sqlcore.CreateAccountParams{
		Email:    email,
		PassHash: passHash,
	})
}

func (a *accounts) GetById(r *http.Request, id uuid.UUID) (sqlcore.Account, error) {
	return a.queries.GetAccountByID(r.Context(), id)
}

func (a *accounts) GetByEmail(r *http.Request, email string) (sqlcore.Account, error) {
	return a.queries.GetAccountByEmail(r.Context(), email)
}

func (a *accounts) UpdatePassword(r *http.Request, id uuid.UUID, passHash string) (sqlcore.Account, error) {
	return a.queries.UpdatePassword(r.Context(), sqlcore.UpdatePasswordParams{
		ID:       id,
		PassHash: passHash,
	})
}

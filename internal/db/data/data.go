package data

import (
	"database/sql"

	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
)

type Databaser struct {
	Accounts   Accounts
	Operations Operations
	Sessions   Sessions
}

func NewDatabaser(db *sql.DB) *Databaser {
	queries := dbcore.New(db)
	return &Databaser{
		Accounts:   NewAccount(queries),
		Operations: NewOperations(queries),
		Sessions:   NewSession(queries),
	}
}

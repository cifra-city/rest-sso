package data

import (
	"database/sql"

	"github.com/cifra-city/rest-sso/internal/db/data/dbcore"
)

type Databaser struct {
	Accounts   Accounts
	Operations Operations
	Sessions   Sessions

	Transaction
}

func NewDBConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewDatabaser(url string) (*Databaser, error) {
	db, err := NewDBConnection(url)
	if err != nil {
		return nil, err
	}
	queries := dbcore.New(db)
	return &Databaser{
		Accounts:    NewAccount(queries),
		Operations:  NewOperations(queries),
		Sessions:    NewSession(queries),
		Transaction: NewTransaction(queries),
	}, nil
}

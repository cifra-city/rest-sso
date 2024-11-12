package handlers

import (
	"net/http"

	"github.com/cifra-city/rest-sso/internal/db/data"
)

type ctxKey int

const (
	SecretUsers ctxKey = iota
)

func Users(r *http.Request) *data.Queries {
	return r.Context().Value(SecretUsers).(*data.Queries)
}

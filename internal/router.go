package service

import (
	"context"

	"github.com/cifra-city/rest-sso/internal/config"

	"github.com/go-chi/chi"
)

func Run(ctx context.Context, cfg config.Config) {
	r := chi.NewRouter()

	r.Route("/cifra-sso", func(r chi.Router) {

	})
}

package service

import (
	"context"
	"net/http"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/handlers"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

func Run(ctx context.Context, cfg config.Config) {
	r := chi.NewRouter()

	r.Route("/cifra-sso", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Route("/auth", func(r chi.Router) {
					r.Post("/register", handlers.Registration)
				})
			})
		})
	})

	logrus.Info("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.Fatal("Server failed to start:", err)
	}
}

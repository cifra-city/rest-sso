package service

import (
	"context"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/handlers"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

func Run(ctx context.Context, cfg config.Config) {
	r := chi.NewRouter()

	// Получаем значение из контекста и приводим его к типу *config.Server
	server, err := cifractx.GetValue[*config.Server](ctx, config.SERVER)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVER, server))

	r.Route("/cifra-sso", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Post("/reg", handlers.Registration)

				r.Route("/auth", func(r chi.Router) {
					r.Post("/login", handlers.Login)
				})
			})
		})
	})

	microServ := httpresp.StartServer(ctx, cfg.Server.Port, r)

	<-ctx.Done()
	httpresp.StopServer(context.Background(), microServ)
}

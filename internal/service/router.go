package service

import (
	"context"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/handlers"
	"github.com/cifra-city/rest-sso/pkg/cifractx"
	"github.com/cifra-city/rest-sso/pkg/httpresp"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context) {
	r := chi.NewRouter()

	service, err := cifractx.GetValue[*config.Service](ctx, config.SERVICE)
	if err != nil {
		logrus.Fatalf("failed to get server from context: %v", err)
	}

	r.Use(cifractx.MiddlewareWithContext(config.SERVICE, service))

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

	server := httpresp.StartServer(ctx, service.Config.Server.Port, r)

	<-ctx.Done()
	httpresp.StopServer(context.Background(), server)
}

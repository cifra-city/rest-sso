package service

import (
	"context"

	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/handlers"
	"github.com/cifra-city/rest-sso/internal/service/middleware"
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
	authMW := middleware.JWTMiddleware(service.Config.JWT.AccessToken.SecretKey, service.Logger)

	r.Route("/cifra-sso", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Post("/registration", handlers.Registration)
				//r.Post("/registration-confirm", handlers.RegistrationConfirm)

				r.Post("/login", handlers.Login)
				//r.Post("/login-confirm", handlers.LoginConfirm)

				r.Route("/user", func(r chi.Router) {
					r.Use(authMW)
					r.Route("/change", func(r chi.Router) {
						//r.Patch("/username", handlers.ChangeUsername)
						//r.Patch("/password", handlers.ChangePassword)
						//r.Patch("/email", handlers.ChangeEmail)
					})
					r.Route("/confirm", func(r chi.Router) {
						//r.Patch("/email", handlers.ConfirmEmail)
						//r.Patch("/reset-password", handlers.ConfirmReestPassword)
						//r.Patch("/reset-username", handlers.ConfirmResetUsername)
						//r.Patch("/reset-email", handlers.ConfirmResetEmail)
					})
					r.Get("/logout", handlers.Logout)
				})

				r.Post("/refresh", handlers.Refresh)
			})
		})
	})

	server := httpresp.StartServer(ctx, service.Config.Server.Port, r)

	<-ctx.Done()
	httpresp.StopServer(context.Background(), server)
}

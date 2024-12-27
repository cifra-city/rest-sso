package service

import (
	"context"
	"time"

	"github.com/cifra-city/cifractx"
	"github.com/cifra-city/httpkit"
	"github.com/cifra-city/rest-sso/internal/config"
	"github.com/cifra-city/rest-sso/internal/service/handlers"
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
	authMW := service.TokenManager.Middleware(service.Config.JWT.AccessToken.SecretKey, service.Logger)
	rateLimiter := httpkit.NewRateLimiter(15, 10*time.Second, 5*time.Minute)
	r.Route("/cifra-sso", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Use(rateLimiter.Middleware)
			r.Route("/public", func(r chi.Router) {
				r.Post("/approve-operation", handlers.ApproveOperation) // approval operation with use email for 15 minutes

				r.Post("/registration-initiate", handlers.RegistrationInitiate) // check if email exists, send code to email
				r.Post("/registration-complete", handlers.RegistrationComplete) // check for approved email address for use
				r.Post("/login-initiate", handlers.LoginInitiate)               // check if email exists, send code to email
				r.Post("/login-complete", handlers.LoginComplete)               // check for approved email address for use
				r.Post("/reset-password-initiate", handlers.ResetPasswordInitiate)
				r.Post("/reset-password-complete", handlers.ResetPasswordComplete)

				r.Post("/refresh", handlers.Refresh)

				r.Route("/user", func(r chi.Router) {
					r.Use(authMW)
					r.Post("/change_username", handlers.ChangeUsername)
					r.Route("/sessions", func(r chi.Router) {
						r.Get("/", handlers.GetUserSessions)
						r.Delete("/", handlers.DeleteSession)
						r.Delete("/terminate", handlers.TerminateSessions)
					})
					r.Post("/logout", handlers.Logout)
				})
				r.Post("/refresh", handlers.Refresh)
			})
		})
	})

	server := httpkit.StartServer(ctx, service.Config.Server.Port, r, service.Logger)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, service.Logger)
}

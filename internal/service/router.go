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
				r.Post("/approve-operation", handlers.ApproveOperation) // approval operation with use email for 15 minutes

				r.Post("/registration", handlers.Registration)                // check if email exists, send code to email
				r.Post("/registration-confirm", handlers.RegistrationConfirm) // check for approved email address for use

				r.Post("/login", handlers.Login)

				r.Post("/forgot-password", handlers.ForgotPassword)
				r.Post("/reset-password", handlers.ResetPassword)

				r.Route("/user", func(r chi.Router) {
					r.Use(authMW)
					//r.Post("/send-code/{operationType}", handlers.SendCode) // user sends code to email with operationType
					r.Route("/change", func(r chi.Router) {
						r.Post("/username", handlers.ChangeUsername) // user sends new username and code
						//r.Post("/password", handlers.ChangePassword)          // user sends new password and code
						//r.Post("/email", handlers.ChangeEmail)                // user sends new email and code
						//r.Post("/email-confirm", handlers.ChangeEmailConfirm) // user sends code to confirm new email
					})

					r.Post("/logout", handlers.Logout)
				})

				r.Post("/refresh", handlers.Refresh)
			})
		})
	})

	server := httpresp.StartServer(ctx, service.Config.Server.Port, r)

	<-ctx.Done()
	httpresp.StopServer(context.Background(), server)
}

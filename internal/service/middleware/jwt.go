package middleware

import (
	"context"
	"net/http"

	"github.com/cifra-city/rest-sso/pkg/cifrajwt"
	"github.com/sirupsen/logrus"
)

// ContextKey defines a type for context keys to avoid collisions.
type ContextKey string

const UserIDKey ContextKey = "userID"

// JWTMiddleware validates the JWT token and injects the user ID into the request context.
func JWTMiddleware(secretKey string, log *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Добавляем запрос в контекст
			ctx := context.WithValue(r.Context(), http.Request{}, r)

			// Извлекаем токен
			tokenString, err := cifrajwt.ExtractToken(ctx)
			if err != nil {
				log.Warnf("Failed to extract token: %v", err)
				http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
				return
			}

			log.Infof("Extracted Token: %s", tokenString)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

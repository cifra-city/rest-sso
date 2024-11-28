package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/cifra-city/rest-sso/pkg/cifrajwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// AuthJWTMiddleware validates the JWT token and injects the user ID into the request context.
func AuthJWTMiddleware(secretKey string, log *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract token from the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				log.Warn("Missing Authorization header")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				log.Warn("Invalid Authorization header format")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			tokenString := parts[1]

			// Verify the JWT token
			claims := &cifrajwt.CustomClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				log.Warn("Invalid or expired token")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Parse user ID from claims
			userID, err := uuid.Parse(claims.Subject)
			if err != nil {
				log.Error("Invalid user ID in token claims")
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Add userID to context and call the next handler
			ctx := context.WithValue(r.Context(), "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

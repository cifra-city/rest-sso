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

type contextKey string

const UserIDKey contextKey = "userID"

// JWTMiddleware validates the JWT token and injects the user ID into the request context.
func JWTMiddleware(secretKey string, log *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

			log.Debugf("Token received: %s", tokenString)

			claims := &cifrajwt.CustomClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err != nil || !token.Valid {
				log.Warnf("Invalid or expired token: %v", err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userID, err := uuid.Parse(claims.Subject)
			if err != nil {
				log.Errorf("Invalid user ID in token claims: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			log.Infof("Claims Subject (UserID): %s", userID)

			// Добавляем userID в контекст
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//sample to use middleware
//userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
//if !ok {
//log.Warn("UserID not found in context")
//httpresp.RenderErr(w, problems.Unauthorized("User not authenticated"))
//return
//}
//logrus.Infof("userID: %v", userID)

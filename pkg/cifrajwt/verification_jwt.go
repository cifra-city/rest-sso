package cifrajwt

import (
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// VerificationJWT validates the JWT token and returns the user ID if valid.
func VerificationJWT(ctx context.Context, log *logrus.Logger, sk string) (uuid.UUID, error) {
	tokenString, err := ExtractToken(ctx)
	if err != nil {
		log.Info("Failed to extract token from context %s", err)
		return uuid.Nil, status.Error(codes.Unauthenticated, "missing or malformed token")
	}

	claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(sk), nil
	})

	if err != nil || !token.Valid {
		log.Info("Token is invalid or expired")
		return uuid.Nil, status.Error(codes.Unauthenticated, "invalid or expired token")
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		log.Info("Failed to parse user ID from claims")
		return uuid.Nil, status.Error(codes.Internal, "invalid user ID format")
	}

	return userID, nil
}

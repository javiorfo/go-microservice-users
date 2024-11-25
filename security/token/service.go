package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-users/config"
)

func Create(permissions map[string][]string, username string) (string, error) {
	return CreateWithDuration(permissions, username, time.Duration(config.TokenDuration * int(time.Second)))
}

func CreateWithDuration(permissions map[string][]string, username string, duration time.Duration) (string, error) {
	tc := config.TokenConfig
	claims := security.TokenClaims{
		Permissions: permissions,
		Audience:    config.TokenAudience,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tc.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Subject:   username,
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tc.SecretKey)
}

func Refresh(oldToken string) (string, error) {
	token, _ := jwt.ParseWithClaims(oldToken, &security.TokenClaims{}, func(token *jwt.Token) (any, error) {
		return config.TokenConfig.SecretKey, nil
	})

	claims, ok := token.Claims.(*security.TokenClaims)
	if !ok {
		return "", errors.New("Invalid token")
	}
	return Create(claims.Permissions, claims.Subject)
}

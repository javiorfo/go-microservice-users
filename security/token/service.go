package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-users/config"
)

func Create(permissions map[string][]string, username string) (string, error) {
	return CreateWithDuration(permissions, username, time.Duration(config.TokenDuration))
}

func CreateWithDuration(permissions map[string][]string, username string, duration time.Duration) (string, error) {
	tc := config.TokenConfig
	claims := security.TokenClaims{
		Username:    username,
		Permissions: permissions,
		Issuer:      tc.Issuer,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tc.SecretKey)
}

func Refresh(oldToken string) (string, error) {
	/* 	claims, err := th.ValidateToken(oldToken)
	   	if err != nil {
	   		return "", err
	   	} */

	// Generate a new token with the same username
	// 	return th.GenerateToken(claims.Username)
	return Create(nil, "Username")
}

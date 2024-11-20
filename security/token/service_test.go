package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/javiorfo/go-microservice-users/config"
	"github.com/stretchr/testify/assert"
)

func TestCreateWithDuration(t *testing.T) {
	// Set up test data
	permissions := map[string][]string{
		"read":  {"users", "posts"},
		"write": {"users", "posts"},
	}
	username := "testuser"
	duration := 1 * time.Hour

	// Call the function being tested
	tokenString, err := CreateWithDuration(permissions, username, duration)
	assert.NoError(t, err)

	// Verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return config.TokenConfig.SecretKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	// Verify the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, username, claims["username"])
	assert.NotEmpty(t, claims["permissions"])
	assert.Equal(t, config.TokenConfig.Issuer, claims["iss"])
	// exp, _ := claims.GetExpirationTime()
	// assert.WithinDuration(t, time.Now().Add(duration), exp.Time, 1*time.Second)
}

package token_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-users/config"
	"github.com/javiorfo/go-microservice-users/security/pwd"
	"github.com/javiorfo/go-microservice-users/security/token"
	"github.com/stretchr/testify/assert"
)

var permission = security.TokenPermission{
	Name:  "read",
	Roles: []string{"user", "admin"},
}

const username = "testuser"
const defaultTokenDuration = 300

func TestCreateWithDuration(t *testing.T) {
	duration := 10 * time.Second

	tokenString, err := token.CreateWithDuration(permission, username, duration)
	assert.NoError(t, err)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return config.TokenConfig.SecretKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, username, claims["sub"])
	assert.NotEmpty(t, claims["permission"])
	assert.Equal(t, config.TokenConfig.Issuer, claims["iss"])
	exp, _ := claims.GetExpirationTime()
	assert.WithinDuration(t, time.Now().Add(duration), exp.Time, 1*time.Second)
}

func TestCreate(t *testing.T) {
	tokenString, err := token.Create(permission, username)
	assert.NoError(t, err)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return config.TokenConfig.SecretKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, username, claims["sub"])
	assert.NotEmpty(t, claims["permission"])
	assert.Equal(t, config.TokenConfig.Issuer, claims["iss"])
	exp, _ := claims.GetExpirationTime()
	assert.WithinDuration(t, time.Now().Add(defaultTokenDuration*time.Second), exp.Time, 1*time.Second)
}

func TestGenerateRandomPassword(t *testing.T) {
	password, err := pwd.GenerateRandomPassword()
	t.Logf("Password generated: %s", password)
	assert.NoError(t, err)
	assert.Len(t, password, 8)
}

func TestGenerateSalt(t *testing.T) {
	salt, err := pwd.GenerateSalt()
	assert.NoError(t, err)
	assert.Len(t, salt, 24)
}

func TestHash(t *testing.T) {
	password := "testPassword"
	salt, err := pwd.GenerateSalt()
	assert.NoError(t, err)

	hash1 := pwd.Hash(password, salt)
	hash2 := pwd.Hash(password, salt)

	assert.Equal(t, hash1, hash2)

	differentSalt, err := pwd.GenerateSalt()
	assert.NoError(t, err)
	hash3 := pwd.Hash(password, differentSalt)
	assert.NotEqual(t, hash1, hash3)
}

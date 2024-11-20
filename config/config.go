package config

import (
	"github.com/javiorfo/go-microservice-lib/env"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-users/internal/database"
)

// Database
var DBDataConnection = database.DBDataConnection{
	Host:        env.GetEnvOr("DB_HOST", "localhost"),
	Port:        env.GetEnvOr("DB_PORT", "5432"),
	DBName:      env.GetEnvOr("DB_NAME", "db_dummy"),
	User:        env.GetEnvOr("DB_USER", "admin"),
	Password:    env.GetEnvOr("DB_PASSWORD", "admin"),
	ShowSQLInfo: true,
}

// App configuration
const AppName = "go-microservice-users"
const AppPort = ":8080"
const AppContextPath = "/users"

// Tracing server configuration
var TracingHost = env.GetEnvOr("TRACING_HOST", "http://localhost:4318")

// Swagger configuration
var SwaggerEnabled = env.GetEnvOr("SWAGGER_ENABLED", true)

// Security
var TokenConfig = security.TokenConfig{
	SecretKey: env.GetEnvOr("JWT_SECRET_KEY", []byte("secret-key")),
	Issuer:    env.GetEnvOr("JWT_ISSUER", "https://users.com"),
	Enabled:   true,
}

var TokenDuration = env.GetEnvOr("JWT_DURATION", 300)

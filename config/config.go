package config

import (
	"github.com/javiorfo/go-microservice-lib/env"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-users/internal/database"
)
// IMPORTANT
// If .env exists it uses the environment variables, otherwise the fallback

// App configuration
const (
	AppName        = "go-microservice-users"
	AppPort        = ":8080"
	AppContextPath = "/users"
	TokenAudience  = "https://something.com.ar"
)

var (
	// Database
	DBDataConnection = database.DBDataConnection{
		Host:        env.GetEnvOr("DB_HOST", "localhost"),
		Port:        env.GetEnvOr("DB_PORT", "5432"),
		DBName:      env.GetEnvOr("DB_NAME", "db_dummy"),
		User:        env.GetEnvOr("DB_USER", "admin"),
		Password:    env.GetEnvOr("DB_PASSWORD", "admin"),
		ShowSQLInfo: true,
	}

	// Tracing server configuration
	TracingHost = env.GetEnvOr("TRACING_HOST", "http://localhost:4318")

	// Swagger configuration
	SwaggerEnabled = env.GetEnvOr("SWAGGER_ENABLED", true)

	// Security
	TokenDuration = env.GetEnvOr("JWT_DURATION", 300)

	TokenConfig = security.TokenConfig{
		SecretKey: env.GetEnvOr("JWT_SECRET_KEY", []byte("secret-key")),
		Issuer:    env.GetEnvOr("JWT_ISSUER", "https://users.com"),
		Enabled:   true,
	}
)

package config

import (
	"github.com/javiorfo/go-microservice-users/internal/database"
	"github.com/javiorfo/go-microservice-lib/env"
)

var DBDataConnection = database.DBDataConnection{
	Host:        env.GetEnvOrDefault("DB_HOST", "localhost"),
	Port:        env.GetEnvOrDefault("DB_PORT", "5432"),
	DBName:      env.GetEnvOrDefault("DB_NAME", "db_dummy"),
	User:        env.GetEnvOrDefault("DB_USER", "admin"),
	Password:    env.GetEnvOrDefault("DB_PASSWORD", "admin"),
	ShowSQLInfo: true,
}

// App configuration
const AppName = "go-microservice-users"
const AppPort = ":8080"
const AppContextPath = "/users"

// Tracing server configuration
var TracingHost = env.GetEnvOrDefault("TRACING_HOST", "http://localhost:4318")

// Swagger configuration
var SwaggerEnabled = env.GetEnvOrDefault("SWAGGER_ENABLED", "true")

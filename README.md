# go-microservice-users
*Go microservice for storing users, credentials, login and registration security*

## Dependencies
Golang, Docker, Make, [Swag tool](https://github.com/swaggo/swag)

## Features
- Architecture
    - Handlers, respositories and services
    - Custom Messages and Errors
    - Pagination and Ordering
    - DB Migrator
- Go 1.23 (at the moment)
- Libraries
    - Web: Fiber
    - ORM: Gorm
    - Security: JWT
    - Validations: Go Playground Validator
    - Unit Test: Testify
    - DB: Postgres
    - Tracing: Opentelemetry
    - Test: Testcontainers
    - OpenAPI: Fiber Swagger
    - Environment: Godot
- Keycloak as Auth Server
- Distributed tracing
    - OpenTelemetry, Micrometer and Jaeger
- Swagger
    - Swaggo & Fiber Swagger
    - Customized with command **make swagger** 
    - Auditory
    - Gorm custom auditory
- Database
    - Postgres for the app
    - Testcontainers for testing

## Files
- [Dockerfile](https://github.com/javiorfo/go-microservice-users/tree/master/Dockerfile)

## Usage
- Executing `make help` all the available commands will be listed. 
- Also the standard Go commands could be used, like `go run main.go`

## Services
- **Create users** POST: /users
- **Get users** GET: /users
- **Get user by ID** GET: /users/{id}
- **Login** POST: /users/login

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
- [Paypal](https://www.paypal.com/donate/?hosted_button_id=FA7SGLSCT2H8G)

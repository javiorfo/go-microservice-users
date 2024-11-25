package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-users/api/handlers"
	"github.com/javiorfo/go-microservice-users/domain/service"
)

const (
	roleUser  = "USERS_USER"
	roleAdmin = "USERS_ADMIN"
)

func User(app fiber.Router, sec security.Securizer, service service.UserService) {
	app.Get("/:id", sec.Secure(roleUser, roleAdmin), handlers.FindById(service))
	app.Get("/", sec.Secure(roleUser, roleAdmin), handlers.FindAll(service))
	app.Post("/", sec.Secure(roleAdmin), handlers.Create(service))
	app.Post("/login", handlers.Login(service))
}

package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice-users/api/routes"
	"github.com/javiorfo/go-microservice-users/config"
	"github.com/javiorfo/go-microservice-users/domain/repository"
	"github.com/javiorfo/go-microservice-users/domain/service"
	"github.com/javiorfo/go-microservice-users/internal/database"
)

func Inject(api fiber.Router) {
	// Database
	db := database.DBinstance

	// Dummy: Repository, Servicer and Routes
	dummyRepository := repository.NewDummyRepository(db)
	dummyService := service.NewDummyService(dummyRepository)
	routes.Dummy(api, config.TokenConfig, dummyService)
}

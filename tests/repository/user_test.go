package user_test 

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/javiorfo/go-microservice-users/domain/model"
	"github.com/javiorfo/go-microservice-users/domain/repository"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var container testcontainers.Container
var repo repository.UserRepository

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	var err error
	container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start container: %s", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		log.Fatalf("Failed to get container host: %s", err)
	}
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("Failed to get container port: %s", err)
	}

	dsn := "host=" + host + " port=" + port.Port() + " user=testuser password=testpass dbname=testdb sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.Permission{}, &model.Role{}); err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}

	repo = repository.NewUserRepository(db)

	// Run the tests
	code := m.Run()

	// Cleanup
	if err := container.Terminate(ctx); err != nil {
		log.Fatalf("Failed to terminate container: %s", err)
	}

	os.Exit(code)
}

func TestUser(t *testing.T) {

	userRecord := model.User{
        Username: "testusername",
        Email: "mail@test.com",
        Password: "p4$$w0rd",
        Salt: "54LT",
        Temporary: false,
        Permission: model.Permission{
            Name: "PERM1",
            Roles: []model.Role{
                {Name: "basic"},
                {Name: "admin"},
            },
        },
    }

	if err := repo.Create(&userRecord); err != nil {
		t.Fatalf("Failed to insert record: %v", err)
	}

	user, err := repo.FindById(userRecord.ID)
	if err != nil {
		t.Fatalf("Failed to query record: %v", err)
	}

	if user.Username != "testusername" {
		t.Errorf("Expected name to be 'testusername', got '%s'", user.Username)
	}

    if user.Permission.Name != "PERM1" {
		t.Errorf("Expected name to be 'PERM1', got '%s'", user.Permission.Name)
	}
}

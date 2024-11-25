package user_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-lib/response"
	"github.com/javiorfo/go-microservice-users/api/request"
	"github.com/javiorfo/go-microservice-users/api/routes"
	"github.com/javiorfo/go-microservice-users/domain/model"
	"github.com/javiorfo/go-microservice-users/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTest() (*fiber.App, *mocks.MockUserService) {
	app := fiber.New()
	mockSec := new(mocks.MockSecurizer)
	mockService := new(mocks.MockUserService)

	routes.User(app, mockSec, mockService)

	return app, mockService
}

// FIND BY ID
func TestFindById(t *testing.T) {
	tests := []struct {
		id           string
		mockReturn   *model.User
		mockError    error
		expectedCode int
	}{
		{"1", &model.User{Username: "username"}, nil, fiber.StatusOK},
		{"2", nil, errors.New("User not found"), fiber.StatusNotFound},
		{"invalid", nil, nil, fiber.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			app, mockService := setupTest()
			if tt.id != "invalid" {
				id, _ := strconv.Atoi(tt.id)
				mockService.On("FindById", uint(id)).Return(tt.mockReturn, tt.mockError)
			}

			req := httptest.NewRequest("GET", "/"+tt.id, nil)
			resp, err := app.Test(req)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, resp.StatusCode)

			if tt.expectedCode == fiber.StatusOK {
				var responseBody model.User
				json.NewDecoder(resp.Body).Decode(&responseBody)
				assert.Equal(t, "username", responseBody.Username)
			}

			mockService.AssertExpectations(t)
		})
	}
}

// FIND ALL
func TestFindAll(t *testing.T) {

	t.Run("Successful", func(t *testing.T) {
		app, mockService := setupTest()
		page := pagination.Page{Page: 1, Size: 10, SortBy: "info", SortOrder: "asc"}
		mockService.On("FindAll", page).Return([]model.User{{ID: 1, Username: "username"}}, nil)

		req := httptest.NewRequest("GET", "/?page=1&size=10&sortBy=info&sortOrder=asc", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var responseBody response.RestResponsePagination[model.User]
		json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.Equal(t, 1, responseBody.Pagination.Total)
		assert.Equal(t, "username", responseBody.Elements[0].Username)

		mockService.AssertExpectations(t)
	})

	t.Run("DB Error", func(t *testing.T) {
		app, mockService := setupTest()
		mockService.On("FindAll", mock.Anything).Return(nil, errors.New("data source error"))

		req := httptest.NewRequest("GET", "/?page=1&size=10&sortBy=id&sortOrder=asc", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})

	t.Run("Pagination Bad Request", func(t *testing.T) {
		app, _ := setupTest()

		req := httptest.NewRequest("GET", "/?page=invalid&size=10&sortBy=id&sortOrder=asc", nil)
		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

// CREATE
func TestCreate(t *testing.T) {

	t.Run("Successful", func(t *testing.T) {
		app, mockService := setupTest()

		userRequest := request.User{
			Username: "username",
			Email:    "mail@mail.com",
			Permission: request.Permission{
				Name:  "PERM",
				Roles: []string{"one", "two"},
			},
		}

        password := "1234"
		mockService.On("Create", mock.Anything).Return(&password, nil)

		body, _ := json.Marshal(userRequest)
		req := httptest.NewRequest("POST", "/", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		var responseBody map[string]string
		json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.Equal(t, password, responseBody["password"])

		mockService.AssertExpectations(t)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		app, _ := setupTest()

		body := `{ "invalid": 10 }`
		req := httptest.NewRequest(fiber.MethodPost, "/", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		app, mockService := setupTest()

		userRequest := request.User{
			Username: "username",
			Email:    "mail@mail.com",
			Permission: request.Permission{
				Name:  "PERM",
				Roles: []string{"one", "two"},
			},
		}
		mockService.On("Create", mock.Anything).Return(nil, errors.New("service error"))

		body, _ := json.Marshal(userRequest)
		req := httptest.NewRequest(fiber.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		mockService.AssertExpectations(t)
	})
}

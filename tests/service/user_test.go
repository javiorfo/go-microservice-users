package user_test

import (
	"errors"
	"testing"

	"github.com/javiorfo/go-microservice-users/domain/model"
	"github.com/javiorfo/go-microservice-users/domain/service"
	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-users/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindUserById(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	id := uint(1)
	expectedUser := &model.User{ID: id}

	mockRepo.On("FindById", id).Return(expectedUser, nil)

	result, err := userService.FindById(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestFindUserByIdNotFound(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	id := uint(1)

	mockRepo.On("FindById", id).Return(nil, errors.New("not found"))

	result, err := userService.FindById(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestFindAllUsers(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	page := pagination.Page{Page: 1, Size: 10, SortBy: "id", SortOrder: "asc"}
	expectedUsers := []model.User{
		{ID: 1},
		{ID: 2},
	}

	mockRepo.On("FindAll", page).Return(expectedUsers, nil)

	result, err := userService.FindAll(page)

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	newUser := &model.User{ID: 1}

	mockRepo.On("Create", newUser).Return(nil)

	randomPassword, err := userService.Create(newUser)

	assert.NoError(t, err)
    assert.Greater(t, len(newUser.Password), 0)
    assert.Equal(t, 8, len(*randomPassword))
    assert.NotEmpty(t, newUser.Salt)
	mockRepo.AssertExpectations(t)
}

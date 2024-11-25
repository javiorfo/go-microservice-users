package mocks

import (
	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-users/domain/model"
	"github.com/stretchr/testify/mock"
)

// Mock Service
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) FindById(id uint) (*model.User, error) {
	args := m.Called(id)
	if user, ok := args.Get(0).(*model.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) FindAll(page pagination.Page) ([]model.User, error) {
	args := m.Called(page)
	if users, ok := args.Get(0).([]model.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) Create(user *model.User) (*string, error) {
	args := m.Called(user)
	if token, ok := args.Get(0).(*string); ok {
		return token, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) Login(username, password string) (string, error) {
	args := m.Called(username, password)
	if token, ok := args.Get(0).(string); ok {
		return token, args.Error(1)
	}
	return "", args.Error(1)
}

// Mock Repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindById(id uint) (*model.User, error) {
	args := m.Called(id)
	if user, ok := args.Get(0).(*model.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	if user, ok := args.Get(0).(*model.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindAll(page pagination.Page) ([]model.User, error) {
	args := m.Called(page)
	if users, ok := args.Get(0).([]model.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

package service

import (
	"errors"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-users/domain/model"
	"github.com/javiorfo/go-microservice-users/domain/repository"
	"github.com/javiorfo/go-microservice-users/security/pwd"
	"github.com/javiorfo/go-microservice-users/security/token"
)

type UserService interface {
	FindById(id uint) (*model.User, error)
	FindAll(pagination.Page) ([]model.User, error)
	Create(*model.User) (*string, error)
	Login(username, password string) (string, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

func (service *userService) FindById(id uint) (*model.User, error) {
	return service.repository.FindById(id)
}

func (service *userService) FindAll(page pagination.Page) ([]model.User, error) {
	return service.repository.FindAll(page)
}

func (service *userService) Create(user *model.User) (*string, error) {
	generatedPassword, err := pwd.GenerateRandomPassword()
	if err != nil {
		return nil, err
	}

	salt, err := pwd.GenerateSalt()
	if err != nil {
		return nil, err
	}
	hashedPassword := pwd.Hash(generatedPassword, salt)

	user.Salt = salt
	user.Password = hashedPassword

    if err := service.repository.Create(user); err != nil {
        return nil, err
    }

    return &generatedPassword, nil
}

func (service *userService) Login(username, password string) (string, error) {
	user, err := service.repository.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if user.VerifyPassword(password) {
		return token.Create(map[string][]string{
            "PERM": {"admin"},
        }, username)
	}
	return "", errors.New("Username or password incorrect")
}

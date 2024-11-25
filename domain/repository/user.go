package repository

import (
	"errors"
	"fmt"

	"github.com/javiorfo/go-microservice-users/domain/model"
	"github.com/javiorfo/go-microservice-lib/pagination"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uint) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindAll(pagination.Page) ([]model.User, error)
	Create(*model.User) error
}

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) FindById(id uint) (*model.User, error) {
	var user model.User

    result := repository.Preload("Permission").Find(&user, "id = ?", id)

	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("User not found")
	}

	return &user, nil
}

func (repository *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User

    result := repository.Preload("Permission").Find(&user, "username = ?", username)

	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("User not found")
	}

	return &user, nil
}

func (repository *userRepository) FindAll(page pagination.Page) ([]model.User, error) {
	var users []model.User

	results := repository.
		Offset(page.Page - 1).
		Limit(page.Size).
		Order(fmt.Sprintf("%s %s", page.SortBy, page.SortOrder)).
		Find(&users)

	if err := results.Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *userRepository) Create(d *model.User) error {
	result := repository.DB.Create(d)

	if err := result.Error; err != nil {
		return fmt.Errorf("Error creating user %v", err)
	}
	return nil
}

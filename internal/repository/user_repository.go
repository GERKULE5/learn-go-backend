// repository/user-repository.go
package repository

import (
	"crud-gin/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)



type UserRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	result := r.db.Create(user)

	if result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}

	return nil
}

func (r *UserRepository) GetById(id int) (*models.User, error) {
	var user models.User
	query := r.db.First(&user, id)

	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}

		return nil, fmt.Errorf("failed to create user: %w", query.Error)
	}

	return &user, nil
}

func (r *UserRepository) Delete(id int)  (*models.User, error) {
	var user models.User
	result := r.db.Delete(&user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return  nil, fmt.Errorf("Failed to delete user with id %d", id)
	}

	return &user, nil
}
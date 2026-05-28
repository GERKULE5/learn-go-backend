// repository/user-repository.go
package repository

import (
	"context"
	app_errors "crud-gin/internal/errors"
	"crud-gin/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id int) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	result := r.db.WithContext(ctx).Create(user)

	if result.Error != nil {
		return fmt.Errorf("Failed to create user: %w", result.Error)
	}

	return nil
}

func (r *UserRepository) GetById(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	query := r.db.WithContext(ctx).First(&user, id)

	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return nil, app_errors.ErrNotFound
		}

		return nil, fmt.Errorf("Failed to get user: %w", query.Error)
	}

	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	query := r.db.WithContext(ctx).Find(&users)

	if query.Error != nil {
		return nil, fmt.Errorf("Failed to get all users: %w", query.Error)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	query := r.db.WithContext(ctx).Save(user)
	if query.Error != nil {
		return fmt.Errorf("Failed to update user: %w", query.Error)
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&models.User{}, id)

	if result.Error != nil {
		// FIX: Delete should only return internal errors when the DB action fails.
		return fmt.Errorf("Failed to delete user with id %d: %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		// FIX: Return not-found when no rows were deleted.
		return app_errors.ErrNotFound
	}

	return nil
}

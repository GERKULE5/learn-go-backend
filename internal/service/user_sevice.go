package service

import (
	"context"
	"crud-gin/internal/dto"
	"crud-gin/internal/models"
	"crud-gin/internal/repository"
)

type UserServiceInterface interface {
	Create(ctx context.Context, data dto.CreateUserDTO) (dto.UserDTO, error)
	GetAll(ctx context.Context) ([]dto.UserDTO, error)
	GetById(ctx context.Context, id int) (dto.UserDTO, error)
	Update(ctx context.Context, id int, data dto.UpdateUserDTO) (dto.UserDTO, error)
	Delete(ctx context.Context, id int) (dto.UserDTO, error)
}

type userService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{repo: repo}
}

func (service *userService) Create(ctx context.Context, data dto.CreateUserDTO) (dto.UserDTO, error) {
	user := &models.User{
		Name:     data.Name,
		Fullname: data.Fullname,
		Age:      data.Age,
		IsSmart:  data.IsSmart,
	}

	if err := service.repo.Create(ctx, user); err != nil {
		return dto.UserDTO{}, err
	}

	return toDTO(user), nil
}

func (service *userService) GetAll(ctx context.Context) ([]dto.UserDTO, error) {
	users, err := service.repo.GetAll(ctx)

	if err != nil {
		return []dto.UserDTO{}, err
	}

	result := make([]dto.UserDTO, 0, len(users))

	for _, user := range users {
		result = append(result, toDTO(user))
	}

	return result, nil
}

func (service *userService) GetById(ctx context.Context, id int) (dto.UserDTO, error) {
	user, err := service.repo.GetById(ctx, id)

	if err != nil {
		return dto.UserDTO{}, err
	}

	return toDTO(user), nil
}

func (service *userService) Update(ctx context.Context, id int, data dto.UpdateUserDTO) (dto.UserDTO, error) {
	user, err := service.repo.GetById(ctx, id)

	if err != nil {
		return dto.UserDTO{}, err
	}

	if data.Name != "" {
		user.Name = data.Name
	}
	if data.Fullname != "" {
		user.Fullname = data.Fullname
	}
	if data.Age != 0 {
		user.Age = data.Age
	}
	if data.IsSmart != nil {
		user.IsSmart = *data.IsSmart
	}

	err = service.repo.Update(ctx, user)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return toDTO(user), nil
}

func (service *userService) Delete(ctx context.Context, id int) (dto.UserDTO, error) {
	user, err := service.repo.GetById(ctx, id)
	if err != nil {
		return dto.UserDTO{}, err
	}

	err = service.repo.Delete(ctx, id)

	if err != nil {
		return dto.UserDTO{}, err
	}

	return toDTO(user), nil
}

func toDTO(u *models.User) dto.UserDTO {
	return dto.UserDTO{
		ID:       int(u.ID),
		Name:     u.Name,
		Fullname: u.Fullname,
		Age:      u.Age,
		IsSmart:  u.IsSmart,
	}
}

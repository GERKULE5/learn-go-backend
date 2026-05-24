// dto/user_dto.go
package dto

import "time"

type CreateUserDTO struct {
	Name        string `json:"name" binding:"required,max=20"`
	Username    string `json:"username" binding:"required,max=16,min=4"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,max=25"`
	Age         int    `json:"age"`
	IsSmart     bool   `json:"is_smart"`
}

type UpdateUserDTO struct {
	Name        string `json:"name" binding:"omitempty,max=20"`
	Username    string `json:"username" binding:"omitempty,max=16,min=4"`
	Email       string `json:"email" binding:"omitempty,email"`
	PhoneNumber string `json:"phone_number" binding:"omitempty,max=25"`
	Age         int    `json:"age"`
	IsSmart     bool   `json:"is_smart"`
}

type UserDTO struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	Age         int        `json:"age"`
	IsSmart     bool       `json:"is_smart"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// dto/user_dto.go
package dto

import "time"

type CreateUserDTO struct {
	Name        string `json:"name" gorm:"not null;size:20"`
	Username    string `json:"username" gorm:"not null;unique;size:16"`
	Email       string `json:"email" gorm:"not null;unique;size:100"`
	PhoneNumber string `json:"phone_number" gorm:"not null;unique;size:25"`
	Age         int    `json:"age" gorm:"default:0"`
	IsSmart     bool   `json:"is_smart" gorm:"default:false"`
}

type UpdateUserDTO struct {
	Name        string `json:"name" gorm:"not null;size:20"`
	Username    string `json:"username" gorm:"not null;unique;size:16"`
	Email       string `json:"email" gorm:"not null;unique;size:100"`
	PhoneNumber string `json:"phone_number" gorm:"not null;unique;size:25"`
	Age         int    `json:"age" gorm:"default:0"`
	IsSmart     bool   `json:"is_smart" gorm:"default:false"`
}

type UserDTO struct {
	ID          int        `json:"id"`
	Name        string     `json:"name" gorm:"not null;size:20"`
	Username    string     `json:"username" gorm:"not null;unique;size:16"`
	Email       string     `json:"email" gorm:"not null;unique;size:100"`
	PhoneNumber string     `json:"phone_number" gorm:"not null;unique;size:25"`
	Age         int        `json:"age" gorm:"default:0"`
	IsSmart     bool       `json:"is_smart" gorm:"default:false"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

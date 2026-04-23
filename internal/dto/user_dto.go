// dto/user_dto.go
package dto

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Fullname string `json:"full_name" binding:"required,min=3,max=200"`
	Age      int    `json:"age" binding:"required,max=150,min=0"`
	IsSmart  bool   `json:"is_smart"`
}

type UpdateUserDTO struct {
	Name     string `json:"name" binding:"min=3,max=100"`
	Fullname string `json:"full_name" binding:"min=3,max=200"`
	Age      int    `json:"age" binding:"omitempty,max=150,min=0"`
	IsSmart  *bool  `json:"is_smart"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Fullname string `json:"full_name"`
	Age      int    `json:"age"`
	IsSmart  bool   `json:"is_smart"`
}

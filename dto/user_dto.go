package dto

type CreateUserDTO struct {
	Name string `json:"name" binding:"required, min=3, max=100`
	Fullname string `json:"full_name" binding:"required, min=3, max=200`
	Age int `json:"age" binding:"required, max=0, min=150`
	IsSmart bool `json:"is_smart"`
}


type UpdateUserDTO struct {
	Name string `json:"name" binding:"min=3, max=100`
	Fullname string `json:"full_name" binding:"min=3, max=200`
	Age int `json:"age" binding:"omitempty, max=0, min=150`
	IsSmart bool `json:"is_smart"`
}

type UserResonse struct {
	Id int `json:"id"` 
	Name string `json:"name"`
	Fullname string `json:"full_name"`
	Age int `json:"age"`
	IsSmart bool `json:"is_smart"`
}

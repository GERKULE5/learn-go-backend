// models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Fullname string `json:"full_name" gorm:"not null"`
	Age int `json:"age"`
	IsSmart bool `json:"is_smart"`
}

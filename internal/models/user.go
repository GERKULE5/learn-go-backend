// models/user.go
package models

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Lastname string `json:"lastname"`
	Age int `json:"age"`
	IsSmart bool `json:"is_smart"`
}

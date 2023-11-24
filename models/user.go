package models


type User struct {
	UserId string `json:"user_id" gorm:"primary_key"`
	Password string `json:"password"`
	Name string `json:"name"`
}
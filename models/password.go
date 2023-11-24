package models

type MyPassword struct {
	Website string `json:"website" gorm:"primary_key"`
	Password string `json:"password"`
	UserId string `json:"user_id"`
}
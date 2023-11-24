package configs

import (
	"github.com/bshome19/go-gin-server/models"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    // DB, err = gorm.Open("sqlite3", "configs/my_passwords.db")
    DB, err = gorm.Open(sqlite.Open("configs/my_passwords.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: "+ err.Error())
    }

    // AutoMigrate will create the necessary tables based on your struct models
    DB.AutoMigrate(&models.MyPassword{}, &models.User{})
}
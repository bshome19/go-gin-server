package configs

import (
	"github.com/bshome19/go-gin-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
    var err error
    DB, err = gorm.Open("sqlite3", "configs/my_passwords.db")
    if err != nil {
        panic("failed to connect database")
    }

    // AutoMigrate will create the necessary tables based on your struct models
    DB.AutoMigrate(&models.MyPasswords{})
}
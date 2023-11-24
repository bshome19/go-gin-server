package handlers

import (
	"net/http"

	"github.com/bshome19/go-gin-server/configs"
	"github.com/bshome19/go-gin-server/models"
	"github.com/bshome19/go-gin-server/utils"
	"github.com/gin-gonic/gin"
)


func GetUserByID(c *gin.Context) {
	var user models.User
	user_id := c.Param("user_id")
	res := configs.DB.Where("user_id = ?", user_id).Find(&user)
	if res.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "No such user"})
		// return user, res.Error
	}
	c.JSON(http.StatusOK, user)
	// return user, nil

}


func CreateUser(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return user, err
	}
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	hashPass, err := utils.HashedPassword(user.Password, salt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user.Password = hashPass
	configs.DB.Create(user)
	c.JSON(http.StatusCreated, user)
	// return *user, nil
}
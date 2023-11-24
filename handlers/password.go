package handlers

import (
	"net/http"
	"github.com/bshome19/go-gin-server/models"
	"github.com/bshome19/go-gin-server/configs"
	"github.com/gin-gonic/gin"
)


func GetPasswords(c *gin.Context) {
	var pb_list []models.MyPassword
	user_id := c.Param("user_id")
	res := configs.DB.Where("user_id = ?", user_id).Find(&pb_list)
	if res.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "No such user"})
		// return user, res.Error
	}
	// configs.DB.Find((&pb_list))
	c.IndentedJSON(http.StatusOK, pb_list)
}

func PostPasswords(c *gin.Context) {
	var pb models.MyPassword

	if err := c.ShouldBindJSON(&pb) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 
	}
	configs.DB.Create(&pb)
	c.IndentedJSON(http.StatusCreated, pb)
}
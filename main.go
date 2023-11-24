package main

import (
	"github.com/bshome19/go-gin-server/configs"
	"github.com/bshome19/go-gin-server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	router.GET("/pbs/:user_id", handlers.GetPasswords)
	router.POST("/pbs", handlers.PostPasswords)

	router.GET("/users/:user_id", handlers.GetUserByID)
	router.POST("/users", handlers.CreateUser)
	router.Run("localhost:8080")
}

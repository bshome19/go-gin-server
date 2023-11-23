package main

import (
	"github.com/gin-gonic/gin"
	"github.com/bshome19/go-gin-server/handlers"
	
)

func main() {
	router := gin.Default()
	router.GET("/pbs", handlers.GetPasswords)
	router.POST("/pbs", handlers.PostPasswords)
	router.Run("localhost:8080")
}
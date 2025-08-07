package main

import (
	"santrikoding/backend-api/config"
	"santrikoding/backend-api/database"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv() // Load environment variables from .env file

	database.InitDB() // Initialize the database connection

	//inisialiasai Gin
	router := gin.Default()

	//membuat route dengan method GET
	router.GET("/", func(ctx *gin.Context) {

		//return response JSON
		ctx.JSON(200, gin.H{
			"message": "Hello Star",
		})
	})

	//mulai server dengan port 3000
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}

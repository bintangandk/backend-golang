package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
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
	router.Run(":3000")
}

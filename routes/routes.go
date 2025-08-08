package routes

import (
	"github.com/gin-gonic/gin"
	"santrikoding/backend-api/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/register", controllers.Register)

	return router

}

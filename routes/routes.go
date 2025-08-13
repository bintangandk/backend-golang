package routes

import (
	"santrikoding/backend-api/controllers"
	"santrikoding/backend-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// Authentication Routes
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	// User Routes
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUsersById)
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)

	// Barang Routes
	router.GET("/api/barang", middlewares.AuthMiddleware(), controllers.FindBarang)
	router.POST("/api/barang", middlewares.AuthMiddleware(), controllers.CreateBarang)
	router.GET("/api/barang/:id", middlewares.AuthMiddleware(), controllers.FindBarangById)
	router.PUT("/api/barang/:id", middlewares.AuthMiddleware(), controllers.UpdateBarang)
	router.DELETE("api/barang/:id", middlewares.AuthMiddleware(), controllers.DeleteBarang)

	return router

}

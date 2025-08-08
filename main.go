package main

import (
	"santrikoding/backend-api/config"
	"santrikoding/backend-api/database"

	"santrikoding/backend-api/routes"
)

func main() {

	config.LoadEnv() // Load environment variables from .env file

	database.InitDB() // Initialize the database connection

	r := routes.SetupRouter() // Setup the router with routes

	//mulai server dengan port 3000
	r.Run(":" + config.GetEnv("APP_PORT", "3000")) // Use the APP_PORT environment variable or default to 3000
}

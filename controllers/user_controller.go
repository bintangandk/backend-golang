package controllers

import (
	"net/http"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {

	// Inisialisasi slice untuk menampung data user
	var Users []models.User

	// Mengambil semua data user dari database
	database.DB.Find(&Users)

	// Kirimkan response sukses dengan data user
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    Users,
	})

}

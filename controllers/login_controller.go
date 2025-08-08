package controllers

import (
	"net/http"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/helpers"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/structs"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login function handles user login requests
func Login(c *gin.Context) {

	// insiasi
	var req = structs.UserLoginRequest{}
	var user = models.User{}

	// validasi input
	// jika input tidak valid, akan mengembalikan error
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Invalid request data",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// mencari user berdasarkan username database
	// jika user tidak ditemukan, akan mengembalikan error
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid username or password",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid username or password",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// jika login berhasil, akan mengembalikan data user dan token
	token := helpers.GenerateToken(user.Username)

	// mengembalikan response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login successful",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			Token:     &token,
		},
	})
}

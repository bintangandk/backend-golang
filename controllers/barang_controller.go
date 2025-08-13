package controllers

import (
	"net/http"
	"santrikoding/backend-api/database"
	"santrikoding/backend-api/helpers"
	"santrikoding/backend-api/models"
	"santrikoding/backend-api/structs"

	"github.com/gin-gonic/gin"
)

func FindBarang(c *gin.Context) {
	// Inisialisasi slice untuk menampung data barang
	var Barangs []models.Barang

	// Mengambil semua data barang dari database
	database.DB.Find(&Barangs)

	// Kirimkan response sukses dengan data barang
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Barang retrieved successfully",
		Data:    Barangs,
	})
}

func CreateBarang(c *gin.Context) {
	var barang models.Barang

	// Bind JSON request body ke struct Barang
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Invalid input data",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Simpan barang ke database
	if err := database.DB.Create(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create barang",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Barang created successfully",
		Data: structs.BarangResponse{
			Id:        barang.Id,
			Nama:      barang.Nama,
			Harga:     barang.Harga,
			Stok:      barang.Stok,
			CreatedAt: barang.CreatedAt.String(),
			UpdatedAt: barang.UpdatedAt.String(),
		},
	})
}

func FindBarangById(c *gin.Context) {
	id := c.Param("id")
	var barang models.Barang

	// Mencari barang berdasarkan ID
	if err := database.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Barang not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Barang retrieved successfully",
		Data: structs.BarangResponse{
			Id:        barang.Id,
			Nama:      barang.Nama,
			Harga:     barang.Harga,
			Stok:      barang.Stok,
			CreatedAt: barang.CreatedAt.String(),
			UpdatedAt: barang.UpdatedAt.String(),
		},
	})
}

func UpdateBarang(c *gin.Context) {
	id := c.Param("id")
	var barang models.Barang

	// Mencari barang berdasarkan ID
	if err := database.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Barang not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Bind JSON request body ke struct Barang
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, structs.ErrorResponse{
			Success: false,
			Message: "Invalid input data",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// Update barang di database
	if err := database.DB.Save(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update barang",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Barang updated successfully",
		Data: structs.BarangResponse{
			Id:        barang.Id,
			Nama:      barang.Nama,
			Harga:     barang.Harga,
			Stok:      barang.Stok,
			CreatedAt: barang.CreatedAt.String(),
			UpdatedAt: barang.UpdatedAt.String(),
		},
	})
}

func DeleteBarang(c *gin.Context) {
	id := c.Param("id")

	var barang models.Barang

	if err := database.DB.First(&barang, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Barang tidak ditemukan",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&barang).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed delete barang",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "barang deleted successfully",
	})
}

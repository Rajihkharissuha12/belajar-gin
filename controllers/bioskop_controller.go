package controllers

import (
	"net/http"

	"tugas13-bioskop/config"
	"tugas13-bioskop/models"

	"github.com/gin-gonic/gin"
)

// ==== CREATE ====

type CreateBioskopInput struct {
	Nama   string  `json:"nama" binding:"required"`
	Lokasi string  `json:"lokasi" binding:"required"`
	Rating float32 `json:"rating"`
}

func CreateBioskop(c *gin.Context) {
	var input CreateBioskopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama dan lokasi tidak boleh kosong"})
		return
	}

	bioskop := models.Bioskop{
		Nama:   input.Nama,
		Lokasi: input.Lokasi,
		Rating: input.Rating,
	}

	if err := config.DB.Create(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menambahkan bioskop"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}


// ==== READ: GET ALL ====

func GetAllBioskop(c *gin.Context) {
	var bioskop []models.Bioskop

	if err := config.DB.Find(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal mendapatkan data bioskop"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}


// ==== READ: GET BY ID ====

func GetBioskopById(c *gin.Context) {
	id := c.Param("id")

	var bioskop models.Bioskop
	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bioskop tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}


// ==== UPDATE ====

type UpdateBioskopInput struct {
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}

func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")

	var bioskop models.Bioskop
	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bioskop tidak ditemukan"})
		return
	}

	var input UpdateBioskopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	// Validasi: Nama & Lokasi tidak boleh kosong jika diisi
	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nama dan lokasi tidak boleh kosong"})
		return
	}

	bioskop.Nama = input.Nama
	bioskop.Lokasi = input.Lokasi
	bioskop.Rating = input.Rating

	config.DB.Save(&bioskop)

	c.JSON(http.StatusOK, bioskop)
}


// ==== DELETE ====

func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")

	var bioskop models.Bioskop
	if err := config.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bioskop tidak ditemukan"})
		return
	}

	config.DB.Delete(&bioskop)

	c.JSON(http.StatusOK, gin.H{"message": "bioskop berhasil dihapus"})
}

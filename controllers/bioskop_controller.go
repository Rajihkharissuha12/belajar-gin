package controllers

import (
	"net/http"
	"tugas13-bioskop/config"
	"tugas13-bioskop/models"

	"github.com/gin-gonic/gin"
)

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

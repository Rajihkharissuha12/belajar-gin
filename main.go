package main

import (
	"tugas13-bioskop/config"
	"tugas13-bioskop/controllers"
	"tugas13-bioskop/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	models.Migrate(config.DB)

	r := gin.Default()
	r.POST("/bioskop", controllers.CreateBioskop)

	r.Run(":8080")
}

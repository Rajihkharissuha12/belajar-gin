package main

import (
	"os"
	"tugas13-bioskop/config"
	"tugas13-bioskop/controllers"
	"tugas13-bioskop/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	models.Migrate(config.DB)

	r := gin.Default()

	// CREATE
	r.POST("/bioskop", controllers.CreateBioskop)

	// READ
	r.GET("/bioskop", controllers.GetAllBioskop)
	r.GET("/bioskop/:id", controllers.GetBioskopById)

	// UPDATE
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)

	// DELETE
	r.DELETE("/bioskop/:id", controllers.DeleteBioskop)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

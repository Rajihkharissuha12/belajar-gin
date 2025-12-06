package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL") // Railway memberikan ini
	if dsn == "" {
		log.Fatal("DATABASE_URL tidak ditemukan")
	}

	// Railway butuh sslmode=require agar connect
	dsn = dsn + " sslmode=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal connect ke database:", err)
	}

	DB = db
	log.Println("Database connected")
}

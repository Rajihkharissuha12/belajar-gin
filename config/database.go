package config

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL tidak ditemukan")
	}

	// Railway biasanya memberi sslmode=disable → harus kita ganti
	// supaya tidak merusak struktur URL
	dsn = strings.Replace(dsn, "sslmode=disable", "sslmode=require", 1)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal connect ke database:", err)
	}

	DB = db
	log.Println("Database connected")
}

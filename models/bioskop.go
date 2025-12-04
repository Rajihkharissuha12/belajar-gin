package models

import "gorm.io/gorm"

type Bioskop struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Bioskop{})
}

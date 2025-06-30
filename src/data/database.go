package data

import (
	"backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {

	db, err := gorm.Open(postgres.Open("host=localhost user=douglas password=douglas dbname=douglas port=5432"), &gorm.Config{})
	if err != nil {
		panic("connection to DB failed!")
	}

	db.AutoMigrate(&models.Bank{}, &models.Card{}, &models.Category{})

	return db
}

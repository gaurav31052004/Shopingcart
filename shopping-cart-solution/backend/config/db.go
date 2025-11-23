package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"shopping-cart/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("shopping-cart.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
	); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	return DB
}

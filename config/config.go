package config

import (
	"ina-gin-crud/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "user:password@tcp(127.0.0.1:3306)/ina?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	log.Println("Database connection established")

	database.AutoMigrate(&models.User{}, &models.Task{})

	DB = database
}

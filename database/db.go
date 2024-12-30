package database

import (
	"log"

	"github.com/matdorneles/vehicle_expenses/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connectionString := "host=localhost user=root password=root dbname=vehicle_expenses port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error connecting to Database!!!")
	}
	DB.AutoMigrate(&models.Vehicle{})
}

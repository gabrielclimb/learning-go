package database

import (
	"log"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	connectionString := "host=localhost user=postgres password=postgres dbname=students port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Fail to connect to DB")
	}
	DB.AutoMigrate(&models.Student{})
}

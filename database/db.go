package database

import (
	"github.com/MikeMwita/person-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=filtronic dbname=edms sslmode=disable password=secret"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	if err := db.AutoMigrate(&models.Person{}); err != nil {
		panic("Failed to auto-migrate database tables: " + err.Error())
	}

	DB = db
	return db
}

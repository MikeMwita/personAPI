//package database
//
//import (
//	"github.com/MikeMwita/person-api/models"
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//)
//
//var (
//	DB *gorm.DB
//)
//
//func ConnectDatabase() *gorm.DB {
//	dsn := "host=localhost user=filtronic dbname=edms sslmode=disable password=secret"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("Failed to connect to the database")
//	}
//	if err := db.AutoMigrate(&models.Person{}); err != nil {
//		panic("Failed to auto-migrate database tables: " + err.Error())
//	}
//
//	DB = db
//	return db
//}

package database

import (
	"fmt"
	"github.com/MikeMwita/person-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

func ConnectDatabase() *gorm.DB {
	// Read environment variables or provide default values for local development
	dbHost := os.Getenv("DATABASE_HOST")
	if dbHost == "" {
		// Set a default value for local development
		dbHost = "filtronicdb"
	}

	dbUser := os.Getenv("DATABASE_USER")
	if dbUser == "" {
		// Set a default value for local development
		dbUser = "filtronic"
	}

	dbPassword := os.Getenv("DATABASE_PASSWORD")
	if dbPassword == "" {
		// Set a default value for local development
		dbPassword = "secret"
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		// Set a default value for local development
		dbName = "edms"
	}

	dbPort := os.Getenv("DATABASE_PORT")
	if dbPort == "" {
		// Set a default value for local development
		dbPort = "5432" // Assuming the default PostgreSQL port is 5432
	}

	// Construct the database connection string
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, dbUser, dbName, dbPassword, dbPort)

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

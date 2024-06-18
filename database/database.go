package database

import (
	"fmt"
	"os"
	"rachanDatingApp/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// InitDB initializes the database connection using environment variables for configuration.
func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string for MySQL
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	// Open a connection to the database
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// AutoMigrate creates/updates the database schema for the given models
	db.AutoMigrate(&models.User{})
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}

package database

import (
	"fmt"
	"log"
	"os"
	"rachanDatingApp/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func InitDB() {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatalf("Error loading .env file: %v", err1)
	}

	// Load environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Printf("DB Host: %s, DB Port: %s, DB User: %s, DB Name: %s\n", dbHost, dbPort, dbUser, dbName)

	// Create the connection string for MySQL
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Printf("Connection String: %s\n", connectionString)

	var err error
	// Open a connection to the database
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// AutoMigrate creates/updates the database schema for the given models
	if err := db.AutoMigrate(&models.User{}, &models.Swipe{}).Error; err != nil {
		panic("failed to migrate database schema: " + err.Error())
	}

	// Check the connection
	if err := db.DB().Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}

	// Connection successful
	fmt.Println("Database connected successfully")
}

func GetDB() *gorm.DB {
	return db
}

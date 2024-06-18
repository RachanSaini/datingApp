package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rachanDatingApp/database"
	"rachanDatingApp/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setup() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize the database
	database.InitDB()

	// Migrate the schema
	database.GetDB().AutoMigrate(&models.User{})
}

func TestLogin(t *testing.T) {
	setup()

	// Create a new Gin router
	router := gin.Default()

	// Define routes
	router.POST("/user/create", CreateUser)
	router.POST("/login", Login)

	// Create a test user
	testUser := models.User{
		Email:    "star@example.com",
		Password: "starpassword",
		Name:     "Test User",
		Gender:   "F",
		DOB:      time.Now(),
	}
	database.GetDB().Create(&testUser)

	// Create a test login request
	loginData := map[string]string{
		"email":    "star@example.com",
		"password": "starpassword",
	}
	loginJSON, _ := json.Marshal(loginData)

	// Create a request to login
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response["token"])
}

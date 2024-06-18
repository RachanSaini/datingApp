package main

import (
	"log"
	"net/http"
	"os"
	"rachanDatingApp/database"
	"rachanDatingApp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	//set gin mode
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)

	//create gin router
	router := gin.Default()
	SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func SetupRoutes(router *gin.Engine) {
	// Health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// User routes
	userGroup := router.Group("/users")
	{
		userGroup.POST("/create", handlers.CreateUser)
		userGroup.POST("/register", handlers.RegisterUser)
	}

	// Authentication
	authgroup := router.Group("/auth")
	{
		authgroup.POST("/login", handlers.Login)
	}

	// Discover
	discoverGroup := router.Group("/discover")
	{
		discoverGroup.GET("/", handlers.DiscoverProfiles)
		discoverGroup.POST("/swipe", handlers.Swipeprofile)
	}
}

package routes

import (
	"net/http"
	"rachanDatingApp/auth"
	"rachanDatingApp/handlers"

	"github.com/gin-gonic/gin"
)

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
	discoverGroup.Use(auth.Authenticate())
	{
		discoverGroup.GET("/", handlers.DiscoverProfiles)
		discoverGroup.POST("/swipe", handlers.SwipeProfile)
	}
}

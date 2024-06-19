package main

import (
	"log"
	"os"
	"rachanDatingApp/database"
	"rachanDatingApp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	//set gin mode
	gin.SetMode(gin.ReleaseMode)

	//create gin router
	router := gin.Default()
	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

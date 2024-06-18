package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"rachanDatingApp/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("SECRET_KEY", "my_secret_key")

	r := gin.New()

	r.POST("/token", func(c *gin.Context) {
		userID, err := strconv.ParseUint(c.PostForm("user_id"), 10, 32)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		token, err := auth.GenerateToken(uint(userID))
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(200, gin.H{"token": token})
	})

	protected := r.Group("/protected")
	protected.Use(auth.Authenticate())
	protected.GET("/", func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.JSON(200, gin.H{"message": "Hello, user " + fmt.Sprintf("%d", userID)})
	})

	log.Fatal(r.Run(":8080"))
}

package handlers

import (
	"net/http"
	"rachanDatingApp/database"
	"rachanDatingApp/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Fetch potential matches for user
func DiscoverProfiles(c *gin.Context) {
	var profiles []models.User

	//current logged in user not discoverable
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not logged in",
		})
		return
	}

	//fetch all profiles except the current logged in user
	database.GetDB().Where("id != ?", userID).Find(&profiles)

	var results []gin.H
	for _, profile := range profiles {
		results = append(results, gin.H{
			"id":     profile.ID,
			"name":   profile.Name,
			"gender": profile.Gender,
			"age":    calculateAge(profile.DOB),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})

}

// Helper function to calculate age from date of birth
func calculateAge(dob time.Time) int {
	today := time.Now()
	age := today.Year() - dob.Year()
	if today.Before(dob.AddDate(age, 0, 0)) {
		age--
	}
	return age
}

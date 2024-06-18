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

	//fetch all users except the current user
	userID := c.GetInt("userId")
	db := database.GetDB()
	db.Not("id", userID).Find(&profiles)

	//exclude the profiles that the user has already swiped
	//future work (this can be implemented by storing the swipe history)

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

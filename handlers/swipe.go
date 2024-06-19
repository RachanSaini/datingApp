package handlers

import (
	"net/http"
	"rachanDatingApp/database"
	"rachanDatingApp/models"

	"github.com/gin-gonic/gin"
)

// SwipeProfile handles the swipe action between users
func SwipeProfile(c *gin.Context) {
	var req struct {
		OtherUserID uint   `json:"otherUserId"`
		Choice      string `json:"choice"` // YES or NO
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetUint("user_id")
	if userId == uint(req.OtherUserID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot swipe on the same user"})
		return
	}

	db := database.GetDB()
	var existingSwipe models.Swipe
	err := db.Where("user_id = ? AND other_user_id = ?", userId, req.OtherUserID).First(&existingSwipe).Error

	if err == nil {
		// Swipe already exists, update it
		existingSwipe.Choice = req.Choice
		db.Save(&existingSwipe)
	} else {
		// Create a new swipe
		existingSwipe = models.Swipe{
			UserID:      uint(userId),
			OtherUserID: req.OtherUserID,
			Choice:      req.Choice,
		}
		db.Create(&existingSwipe)
	}

	// Check for a match
	if req.Choice == "YES" {
		var otherSwipe models.Swipe
		if err := db.Where("user_id = ? AND other_user_id = ? AND choice = ?", req.OtherUserID, userId, "YES").First(&otherSwipe).Error; err == nil {
			// Create a match
			match := models.Match{
				UserID1: uint(userId),
				UserID2: req.OtherUserID,
			}
			db.Create(&match)
			c.JSON(http.StatusOK, gin.H{"results": gin.H{"matched": true, "matchID": match.ID}})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"results": gin.H{"matched": false, "matchID": existingSwipe.ID}})
}

package handlers

import (
	"net/http"
	"rachanDatingApp/database"
	"rachanDatingApp/models"

	"github.com/gin-gonic/gin"
)

func Swipeprofile(c *gin.Context) {
	var req struct {
		OtherUserID uint   `json:"otherUserId"`
		Choice      string `json:"choice"` // YES or NO
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set at authentication
	userId := c.GetInt("userID")

	//checking existing swipe
	var existingSwipe models.Swipe
	db := database.GetDB()
	if err := db.Where("user_id = ? AND other_user_id = ?", userId, req.OtherUserID).First(&existingSwipe).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Swipe already exists"})
		return
	}

	// Create a new swipe record
	newSwipe := models.Swipe{
		UserID:      uint(userId),
		OtherUserID: req.OtherUserID,
		Choice:      req.Choice,
	}
	if req.Choice == "YES" {
		//check other user swiped yes as well
		var otherSwipe models.Swipe
		if err := db.Where("user_id = ? AND other_user_id = ? AND choice = ?", req.OtherUserID, userId, "YES").First(&otherSwipe).Error; err == nil {
			newSwipe.Matched = true
			otherSwipe.Matched = true
			db.Save(&otherSwipe)
		}
	}

	db.Create(&newSwipe)

	c.JSON(http.StatusOK, gin.H{"results": gin.H{"matched": newSwipe.Matched, "matchID": newSwipe.ID}})

}

package models

import "gorm.io/gorm"

// Match represents a match between two users
type Match struct {
	gorm.Model
	UserID1 uint `gorm:"not null"`
	UserID2 uint `gorm:"not null"`
}

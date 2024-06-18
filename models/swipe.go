package models

import "time"

type Swipe struct {
	ID          uint   `gorm:"primary_key"`
	UserID      uint   `gorm:"not null"`
	OtherUserID uint   `gorm:"not null"`
	Choice      string `gorm:"not null"` // YES or NO
	Matched     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

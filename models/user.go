package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Gender    string    `gorm:"not null"`
	DOB       time.Time `gorm:"not null"`
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

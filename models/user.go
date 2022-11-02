package models

import "time"

type User struct {
	UID          int    `gorm:"primaryKey;not null;column:uid"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	ProfileImage string
	Verified     bool      `gorm:"default:false;not null"`
	CreatedAt    time.Time `gorm:"not null"`
	Version      int       `gorm:"default:0;not null"`
}

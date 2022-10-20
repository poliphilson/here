package models

type User struct {
	UID          int    `gorm:"primaryKey;not null;column:uid"`
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Name         string `gorm:"unique;not null"`
	ProfileImage []byte
	Verified     bool `gorm:"default:false;not null"`
	Version      int  `gorm:"default:0;not null"`
}

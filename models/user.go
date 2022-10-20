package models

type User struct {
	UID          int    `gorm:"primaryKey;not null;column:uid"`
	Id           string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Name         string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	ProfileImage []byte
	Version      int `gorm:"default:0;not null"`
}

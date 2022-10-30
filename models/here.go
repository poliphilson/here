package models

import (
	"time"

	datatype "github.com/poliphilson/here/datatype"
)

type Here struct {
	HID       int               `gorm:"primaryKey;not null;column:hid"`
	Uid       int               `gorm:"not null"`
	User      User              `gorm:"foreignKey:Uid"`
	CreatedAt time.Time         `gorm:"not null"`
	Contents  string            `gorm:"not null"`
	Delete    bool              `gorm:"default:false;not null"`
	Image     bool              `gorm:"not null"`
	Video     bool              `gorm:"not null"`
	Location  datatype.Location `gorm:"not null"`
	Version   int               `gorm:"default:0;not null"`
}

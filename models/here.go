package models

import (
	"time"

	datatype "github.com/poliphilson/here/data_type"
)

type Here struct {
	HID      int       `gorm:"primaryKey;not null;column:hid"`
	Uid      int       `gorm:"not null"`
	User     User      `gorm:"foreignKey:Uid"`
	Date     time.Time `gorm:"not null"`
	Contents string
	Image1   []byte
	Image2   []byte
	Image3   []byte
	Delete   bool              `gorm:"default:false;not null"`
	Location datatype.Location `gorm:"not null"`
	Version  int               `gorm:"default:0;not null"`
}

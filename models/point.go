package models

import (
	"time"

	datatype "github.com/poliphilson/here/datatype"
)

type Point struct {
	PID         int  `gorm:"primaryKey;not null;column:pid"`
	Uid         int  `gorm:"not null"`
	User        User `gorm:"foreignKey:Uid"`
	Description string
	CreatedAt   time.Time         `gorm:"not null"`
	Location    datatype.Location `gorm:"not null"`
	Delete      bool              `gorm:"default:false;not null"`
	Version     int               `gorm:"default:0;not null"`
}

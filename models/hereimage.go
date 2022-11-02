package models

type HereImage struct {
	Hid   int    `gorm:"not null"`
	Here  Here   `gorm:"foreignKey:Hid"`
	Image string `gorm:"unique;not null"`
}

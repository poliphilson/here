package models

type HereVideo struct {
	Hid   int    `gorm:"not null"`
	Here  Here   `gorm:"foreignKey:Hid"`
	Video string `gorm:"unique;not null"`
}

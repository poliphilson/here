package models

type HereAddress struct {
	Hid             int  `gorm:"not null"`
	Here            Here `gorm:"foreignKey:Hid"`
	Name            string
	Street          string
	Country         string
	AdminArea       string
	SubArea         string
	Locality        string
	SubLocality     string
	Thoroughfare    string
	SubThoroughfare string
}

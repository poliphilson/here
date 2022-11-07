package response

import (
	"time"

	"github.com/poliphilson/here/datatype"
)

type SimpleHere struct {
	Hid        int               `json:"hid"`
	CreatedAt  time.Time         `json:"created_at"`
	Contents   string            `json:"contents"`
	Location   datatype.Location `json:"location"`
	Image      bool              `json:"image"`
	Video      bool              `json:"video"`
	IsPrivated bool              `json:"is_privated"`
}

type DetailHere struct {
	Here   SimpleHere `json:"here"`
	Images []string   `json:"images"`
	Videos []string   `json:"videos"`
}

type Point struct {
	Pid         int               `json:"pid"`
	CreatedAt   time.Time         `json:"created_at"`
	Description string            `json:"description"`
	Location    datatype.Location `json:"location"`
}

type User struct {
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
}

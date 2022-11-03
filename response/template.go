package response

import (
	"time"

	"github.com/poliphilson/here/datatype"
)

type SimpleHere struct {
	Hid       int               `json:"hid"`
	CreatedAt time.Time         `json:"created_at"`
	Contents  string            `json:"contents"`
	Location  datatype.Location `json:"location"`
	Image     bool              `json:"image"`
	Video     bool              `json:"video"`
}

type DetailHere struct {
	Here   SimpleHere `json:"here"`
	Images []string   `json:"images"`
	Videos []string   `json:"videos"`
}

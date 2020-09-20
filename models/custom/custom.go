package custom

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS  string         `xml:"xmlns,attr,omitempty"`
	Custom *models.Custom `xml:"custom,omitempty"`
}

package custom

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	Custom *models.Custom `xml:"custom,omitempty"`
}

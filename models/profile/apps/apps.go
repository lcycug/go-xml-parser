package apps

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	Apps []*models.ApplicationVisibilities `xml:"applicationVisibilities,omitempty"`
}

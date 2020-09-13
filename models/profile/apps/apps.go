package apps

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS string                            `xml:"xmlns,attr,omitempty"`
	Apps  []*models.ApplicationVisibilities `xml:"applicationVisibilities,omitempty"`
}

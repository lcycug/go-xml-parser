package tabs

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	Tabs []*models.TabVisibilities `xml:"tabVisibilities,omitempty"`
}

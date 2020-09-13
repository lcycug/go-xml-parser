package tabs

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS string                    `xml:"xmlns,attr,omitempty"`
	Tabs  []*models.TabVisibilities `xml:"tabVisibilities,omitempty"`
}

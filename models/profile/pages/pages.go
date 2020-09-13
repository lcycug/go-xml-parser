package pages

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS string                 `xml:"xmlns,attr,omitempty"`
	Pages []*models.PageAccesses `xml:"pageAccesses,omitempty"`
}

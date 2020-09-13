package pages

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	Pages []*models.PageAccesses `xml:"pageAccesses,omitempty"`
}

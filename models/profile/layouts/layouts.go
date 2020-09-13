package layouts

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS   string                      `xml:"xmlns,attr,omitempty"`
	Layouts []*models.LayoutAssignments `xml:"layoutAssignments,omitempty"`
}

package layouts

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	Layouts []*models.LayoutAssignments `xml:"layoutAssignments,omitempty"`
}

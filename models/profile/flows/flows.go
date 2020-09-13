package flows

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	Flows []*models.FlowAccesses `xml:"flowAccesses,omitempty"`
}

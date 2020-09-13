package flows

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS string                 `xml:"xmlns,attr,omitempty"`
	Flows []*models.FlowAccesses `xml:"flowAccesses,omitempty"`
}

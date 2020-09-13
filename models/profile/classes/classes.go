package classes

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS   string                  `xml:"xmlns,attr,omitempty"`
	Classes []*models.ClassAccesses `xml:"classAccesses,omitempty"`
}

package classes

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	Classes []*models.ClassAccesses `xml:"classAccesses,omitempty"`
}

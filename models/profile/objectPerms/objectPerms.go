package objectPerms

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS       string                      `xml:"xmlns,attr,omitempty"`
	ObjectPerms []*models.ObjectPermissions `xml:"objectPermissions,omitempty"`
}

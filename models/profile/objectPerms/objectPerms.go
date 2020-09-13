package objectPerms

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	ObjectPerms []*models.ObjectPermissions `xml:"objectPermissions,omitempty"`
}

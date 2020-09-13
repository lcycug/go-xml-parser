package fieldPerms

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS      string                     `xml:"xmlns,attr,omitempty"`
	FieldPerms []*models.FieldPermissions `xml:"fieldPermissions,omitempty"`
}

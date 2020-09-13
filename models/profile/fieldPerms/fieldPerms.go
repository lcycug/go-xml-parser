package fieldPerms

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	FieldPerms []*models.FieldPermissions `xml:"fieldPermissions,omitempty"`
}

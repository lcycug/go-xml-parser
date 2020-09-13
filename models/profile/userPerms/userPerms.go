package userPerms

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	UserPerms []*models.UserPermissions `xml:"userPermissions,omitempty"`
}

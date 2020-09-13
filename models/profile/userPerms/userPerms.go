package userPerms

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS     string                    `xml:"xmlns,attr,omitempty"`
	UserPerms []*models.UserPermissions `xml:"userPermissions,omitempty"`
}

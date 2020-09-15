package userPerms

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS     string `xml:"xmlns,attr,omitempty"`
	UserPerms `xml:"userPermissions,omitempty"`
}

type UserPerms []*models.UserPermissions

//Implement Sort interface
func (p UserPerms) Len() int      { return len(p) }
func (p UserPerms) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

//ByName implements sort.Interface by providing Less and using the Len and
//Swap methods of the embedded UserPerms value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.UserPerms[i].Name,
		n.Profile.UserPerms[j].Name) < 0
}

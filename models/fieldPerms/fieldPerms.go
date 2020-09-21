package fieldPerms

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS      string `xml:"xmlns,attr,omitempty"`
	FieldPerms `xml:"fieldPermissions,omitempty"`
}

type FieldPerms []*models.FieldPermissions

// Implement Sort interface
func (p FieldPerms) Len() int      { return len(p) }
func (p FieldPerms) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded FieldPerms value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.FieldPerms[i].Field,
		n.Profile.FieldPerms[j].Field) < 0
}

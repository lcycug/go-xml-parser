package objectPerms

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS       string `xml:"xmlns,attr,omitempty"`
	ObjectPerms `xml:"objectPermissions,omitempty"`
}

type ObjectPerms []*models.ObjectPermissions

// Implement Sort interface
func (p ObjectPerms) Len() int      { return len(p) }
func (p ObjectPerms) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded ObjectPerms value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.ObjectPerms[i].Object,
		n.Profile.ObjectPerms[j].Object) < 0
}

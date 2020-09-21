package classes

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS   string `xml:"xmlns,attr,omitempty"`
	Classes `xml:"classAccesses,omitempty"`
}

type Classes []*models.ClassAccesses

// Implement Sort interface
func (c Classes) Len() int      { return len(c) }
func (c Classes) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Classes value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.Classes[i].ApexClass,
		n.Profile.Classes[j].ApexClass) < 0
}

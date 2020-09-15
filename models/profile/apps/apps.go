package apps

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS string `xml:"xmlns,attr,omitempty"`
	Apps  `xml:"applicationVisibilities,omitempty"`
}

type Apps []*models.ApplicationVisibilities

// Implement Sort interface
func (a Apps) Len() int      { return len(a) }
func (a Apps) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Apps value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.Apps[i].Name, n.Profile.Apps[j].Name) < 0
}

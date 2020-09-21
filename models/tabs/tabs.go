package tabs

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS string `xml:"xmlns,attr,omitempty"`
	Tabs  `xml:"tabVisibilities,omitempty"`
}

type Tabs []*models.TabVisibilities

// Implement Sort interface
func (p Tabs) Len() int      { return len(p) }
func (p Tabs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Tabs value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.Tabs[i].Tab, n.Profile.Tabs[j].Tab) < 0
}

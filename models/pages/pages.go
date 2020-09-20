package pages

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS string `xml:"xmlns,attr,omitempty"`
	Pages `xml:"pageAccesses,omitempty"`
}

type Pages []*models.PageAccesses

// Implement Sort interface
func (p Pages) Len() int      { return len(p) }
func (p Pages) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Pages value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.Pages[i].Name, n.Profile.Pages[j].Name) < 0
}

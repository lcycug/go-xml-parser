package layouts

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS   string `xml:"xmlns,attr,omitempty"`
	Layouts `xml:"layoutAssignments,omitempty"`
}

type Layouts []*models.LayoutAssignments

// Implement Sort interface
func (l Layouts) Len() int      { return len(l) }
func (l Layouts) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Layouts value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.Layouts[i].Name,
		n.Profile.Layouts[j].Name) < 0
}

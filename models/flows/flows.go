package flows

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS string `xml:"xmlns,attr,omitempty"`
	Flows `xml:"flowAccesses,omitempty"`
}

type Flows []*models.FlowAccesses

// Implement Sort interface
func (f Flows) Len() int      { return len(f) }
func (f Flows) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Flows value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.Flows[i].Name, n.Profile.Flows[j].Name) < 0
}

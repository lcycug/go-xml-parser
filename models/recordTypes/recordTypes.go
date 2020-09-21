package recordTypes

import (
	"strings"

	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS       string `xml:"xmlns,attr,omitempty"`
	RecordTypes `xml:"recordTypeVisibilities,omitempty"`
}

type RecordTypes []*models.RecordTypeVisibilities

// Implement Sort interface
func (p RecordTypes) Len() int      { return len(p) }
func (p RecordTypes) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded RecordTypes value.
type ByName struct{ Profile }

func (n ByName) Less(i, j int) bool {
	return strings.Compare(n.Profile.RecordTypes[i].RecordType,
		n.Profile.RecordTypes[j].RecordType) < 0
}

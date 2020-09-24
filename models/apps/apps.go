package apps

import (
	"strings"

	"github.com/fatih/structs"
	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS string                            `xml:"xmlns,attr,omitempty"`
	Apps  []*models.ApplicationVisibilities `xml:"applicationVisibilities,omitempty"`
}

// Profile implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Apps value.
func (p Profile) Len() int      { return len(p.Apps) }
func (p Profile) Swap(i, j int) { p.Apps[i], p.Apps[j] = p.Apps[j], p.Apps[i] }
func (p Profile) Less(i, j int) bool {
	return strings.Compare(p.Apps[i].Application, p.Apps[j].Application) < 0
}

// Get the first element of the Profile.
func (p Profile) First() interface{} {
	if p.Len() > 0 {
		return p.Apps[0]
	} else {
		return nil
	}
}

// Get the Name of this type of component as the tab name of the excel workbook.
func (p Profile) Name() string {
	return structs.Name(p.First())
}

// Get all values mapped with fields.
func (p Profile) Values() map[*structs.Field][]interface{} {
	var v = make(map[*structs.Field][]interface{})
	for i, f := range structs.Fields(p.First()) {
		for _, a := range p.Apps {
			v[f] = append(v[f], structs.Values(a)[i])
		}
	}
	return v
}

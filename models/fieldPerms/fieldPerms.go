package fieldPerms

import (
	"strings"

	"github.com/fatih/structs"
	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS      string                     `xml:"xmlns,attr,omitempty"`
	FieldPerms []*models.FieldPermissions `xml:"fieldPermissions,omitempty"`
}

// Profile implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded FieldPerms value.
func (p Profile) Len() int { return len(p.FieldPerms) }
func (p Profile) Less(i, j int) bool {
	return strings.Compare(p.FieldPerms[i].Field, p.FieldPerms[j].Field) < 0
}
func (p Profile) Swap(i, j int) {
	p.FieldPerms[i], p.FieldPerms[j] = p.FieldPerms[j], p.FieldPerms[i]
}

// This method is used to get the first element of the Profile.
func (p Profile) First() interface{} {
	if p.Len() > 0 {
		return p.FieldPerms[0]
	} else {
		return nil
	}
}

// This method is used to get the Name of this component as the tab name of
// the excel workbook.
func (p Profile) Name() string {
	return structs.Name(p.First())
}

// Get all values mapped with fields.
func (p Profile) Values() map[*structs.Field][]interface{} {
	var v = make(map[*structs.Field][]interface{})
	for i, f := range structs.Fields(p.First()) {
		for _, a := range p.FieldPerms {
			v[f] = append(v[f], structs.Values(a)[i])
		}
	}
	return v
}

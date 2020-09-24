package userPerms

import (
	"strings"

	"github.com/fatih/structs"
	"github.com/lcycug/go-xml-parser/models"
)

type Profile struct {
	XMLNS     string                    `xml:"xmlns,attr,omitempty"`
	UserPerms []*models.UserPermissions `xml:"userPermissions,omitempty"`
}

// Profile implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded UserPerms value.
func (p Profile) Len() int { return len(p.UserPerms) }
func (p Profile) Swap(i, j int) {
	p.UserPerms[i], p.UserPerms[j] = p.UserPerms[j], p.UserPerms[i]
}
func (p Profile) Less(i, j int) bool {
	return strings.Compare(p.UserPerms[i].Name, p.UserPerms[j].Name) < 0
}

// This method is used to get the Name of this component as the tab name of
// the excel workbook.
func (p Profile) Name() string {
	return structs.Name(p.First())
}

// This method is used to get the first element of the Profile.
func (p Profile) First() interface{} {
	if p.Len() > 0 {
		return p.UserPerms[0]
	} else {
		return nil
	}
}

// Get all values mapped with fields.
func (p Profile) Values() map[*structs.Field][]interface{} {
	var v = make(map[*structs.Field][]interface{})
	for i, f := range structs.Fields(p.First()) {
		for _, a := range p.UserPerms {
			v[f] = append(v[f], structs.Values(a)[i])
		}
	}
	return v
}

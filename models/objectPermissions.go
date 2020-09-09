package models

import "strings"

// ObjectPermissions is used to store a list of permissions for a SObject.
type ObjectPermissions struct {
	AllowCreate      bool   `xml:"allowCreate"`
	AllowDelete      bool   `xml:"allowDelete"`
	AllowEdit        bool   `xml:"allowEdit"`
	AllowRead        bool   `xml:"allowRead"`
	ModifyAllRecords bool   `xml:"modifyAllRecords"`
	ViewAllRecords   bool   `xml:"viewAllRecords"`
	Object           string `xml:"object"`
}

type ObjectPerms []*ObjectPermissions

// Implement Sort interface
func (p ObjectPerms) Len() int      { return len(p) }
func (p ObjectPerms) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByUserPermName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded ObjectPerms value.
type ByObjectName struct{ ObjectPerms }

func (n ByObjectName) Less(i, j int) bool {
	return strings.Compare(n.ObjectPerms[i].Object, n.ObjectPerms[j].Object) < 0
}

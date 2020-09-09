package models

import "strings"

// FieldPermissions is used to store field permissions for this Profile.
type FieldPermissions struct {
	Editable bool   `xml:"editable"`
	Field    string `xml:"field"`
	Readable bool   `xml:"readable"`
}

type FieldPerms []*FieldPermissions

// Implement Sort interface
func (p FieldPerms) Len() int      { return len(p) }
func (p FieldPerms) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByUserPermName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded FieldPerms value.
type ByFieldName struct{ FieldPerms }

func (n ByFieldName) Less(i, j int) bool {
	return strings.Compare(n.FieldPerms[i].Field, n.FieldPerms[j].Field) < 0
}

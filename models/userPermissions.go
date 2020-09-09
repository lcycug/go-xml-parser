package models

import (
	"fmt"
	"strings"
)

// UserPermissions is used to store User Permissions specified for the Profile.
type UserPermissions struct {
	Name    string `xml:"name"`
	Enabled bool   `xml:"enabled"`
}

// This method is used to return a formatted string.
func (p UserPermissions) String() string {
	return fmt.Sprintf("Name: %s, Enabled: %t", p.Name, p.Enabled)
}

type UserPerms []*UserPermissions

// Implement Sort interface
func (p UserPerms) Len() int      { return len(p) }
func (p UserPerms) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByUserPermName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded UserPerms value.
type ByUserPermName struct{ UserPerms }

func (n ByUserPermName) Less(i, j int) bool {
	return strings.Compare(n.UserPerms[i].Name, n.UserPerms[j].Name) < 0
}

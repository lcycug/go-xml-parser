package models

import "fmt"

// UserPermissions is used to store User Permissions specified for the Profile.
type UserPermissions struct {
	Name    string `xml:"name"`
	Enabled bool   `xml:"enabled"`
}

// This method is used to return a formatted string.
func (p UserPermissions) String() string {
	return fmt.Sprintf("Name: %s, Enabled: %t", p.Name, p.Enabled)
}

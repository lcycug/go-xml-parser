package models

import "fmt"

type UserPermissions struct {
	Name    string `xml:"name"`
	Enabled bool   `xml:"enabled"`
}

func (p UserPermissions) String() string {
	return fmt.Sprintf("Name: %s, Enabled: %t", p.Name, p.Enabled)
}

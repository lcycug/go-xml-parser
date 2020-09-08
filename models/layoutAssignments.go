package models

import "fmt"

// LayoutAssignments is used to tell if a Layout is specified for a Profile
type LayoutAssignments struct {
	Layout string `xml:"layout"`
}

// This method is used to return a formatted string.
func (l LayoutAssignments) String() string {
	return fmt.Sprintf("Layout: %s", l.Layout)
}

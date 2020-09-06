package models

import "fmt"

type LayoutAssignments struct {
	Layout string `xml:"layout"`
}

func (l LayoutAssignments) String() string {
	return fmt.Sprintf("Layout: %s", l.Layout)
}

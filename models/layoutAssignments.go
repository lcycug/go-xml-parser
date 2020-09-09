package models

import (
	"fmt"
	"strings"
)

// LayoutAssignments is used to tell if a Layout is specified for a Profile
type LayoutAssignments struct {
	Layout string `xml:"layout"`
}

// This method is used to return a formatted string.
func (l LayoutAssignments) String() string {
	return fmt.Sprintf("Layout: %s", l.Layout)
}

type Layouts []*LayoutAssignments

// Implement Sort interface
func (l Layouts) Len() int      { return len(l) }
func (l Layouts) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

// ByUserPermName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Layouts value.
type ByLayoutName struct{ Layouts }

func (n ByLayoutName) Less(i, j int) bool {
	return strings.Compare(n.Layouts[i].Layout, n.Layouts[j].Layout) < 0
}

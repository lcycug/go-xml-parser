package models

import (
	"fmt"
	"strings"
)

// TabVisibilities is used to store visibilities for Tabs
type TabVisibilities struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}

// This method is used to return a formatted string.
func (t TabVisibilities) String() string {
	return fmt.Sprintf("Tab: %s, Visibility: %s", t.Tab, t.Visibility)
}

type Tabs []*TabVisibilities

// Implement Sort interface
func (p Tabs) Len() int      { return len(p) }
func (p Tabs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByUserPermName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Tabs value.
type ByTabName struct{ Tabs }

func (n ByTabName) Less(i, j int) bool {
	return strings.Compare(n.Tabs[i].Tab, n.Tabs[j].Tab) < 0
}

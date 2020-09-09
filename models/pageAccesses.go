package models

import "strings"

// PageAccesses is used to specify the accessibility of a Visualforce Page.
type PageAccesses struct {
	ApexPage string `xml:"apexPage"`
	Enabled  bool   `xml:"enabled"`
}

type Pages []*PageAccesses

// Implement Sort interface
func (p Pages) Len() int      { return len(p) }
func (p Pages) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByUserPermName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Pages value.
type ByPageName struct{ Pages }

func (n ByPageName) Less(i, j int) bool {
	return strings.Compare(n.Pages[i].ApexPage, n.Pages[j].ApexPage) < 0
}

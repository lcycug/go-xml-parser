package models

import (
	"fmt"
	"strings"
)

// ApplicationVisibilities is used to tell if an application is visible and defaulted to a Profile.
type ApplicationVisibilities struct {
	Application string `xml:"application"`
	Default     bool   `xml:"default"`
	Visible     bool   `xml:"visible"`
}

// This method is used to return a formatted string.
func (a ApplicationVisibilities) String() string {
	return fmt.Sprintf("Application: %s, Default: %t, Visible: %t", a.Application, a.Default, a.Visible)
}

type Apps []*ApplicationVisibilities

// Implement Sort interface
func (a Apps) Len() int      { return len(a) }
func (a Apps) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// ByAppName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Apps value.
type ByAppName struct{ Apps }

func (n ByAppName) Less(i, j int) bool {
	return strings.Compare(n.Apps[i].Application, n.Apps[j].Application) < 0
}

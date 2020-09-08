package models

import "fmt"

// TabVisibilities is used to store visibilities for Tabs
type TabVisibilities struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}

// This method is used to return a formatted string.
func (t TabVisibilities) String() string {
	return fmt.Sprintf("Tab: %s, Visibility: %s", t.Tab, t.Visibility)
}

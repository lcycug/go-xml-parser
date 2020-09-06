package models

import "fmt"

type TabVisibilities struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}

func (t TabVisibilities) String() string {
	return fmt.Sprintf("Tab: %s, Visibility: %s", t.Tab, t.Visibility)
}

package models

import "fmt"

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

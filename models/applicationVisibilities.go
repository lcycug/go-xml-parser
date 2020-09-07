package models

import "fmt"

type ApplicationVisibilities struct {
	Application string `xml:"application"`
	Default     bool   `xml:"default"`
	Visible     bool   `xml:"visible"`
}

func (a ApplicationVisibilities) String() string {
	return fmt.Sprintf("Application: %s, Default: %t, Visible: %t", a.Application, a.Default, a.Visible)
}

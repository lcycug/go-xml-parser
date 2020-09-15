package models

//ApplicationVisibilities is used to tell if an application is visible and
//defaulted to a Profile.
type ApplicationVisibilities struct {
	Name    string `xml:"application"`
	Default bool   `xml:"default"`
	Visible bool   `xml:"visible"`
}

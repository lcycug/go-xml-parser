package models

type ApplicationVisibilities struct {
	Application string `xml:"application"`
	Default     bool   `xml:"default"`
	Visible     bool   `xml:"visible"`
}

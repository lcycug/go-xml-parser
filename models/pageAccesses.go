package models

// PageAccesses is used to specify the accessibility of a Visualforce Page.
type PageAccesses struct {
	ApexPage string `xml:"apexPage"`
	Enabled  bool   `xml:"enabled"`
}

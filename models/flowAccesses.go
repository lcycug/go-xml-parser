package models

// FlowAccesses is used to tell if a flow is visible to a Profile
type FlowAccesses struct {
	Name    string `xml:"flow"`
	Enabled bool   `xml:"enabled"`
}

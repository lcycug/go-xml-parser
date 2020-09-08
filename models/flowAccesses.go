package models

// FlowAccesses is used to tell if a flow is visible to a Profile
type FlowAccesses struct {
	Enabled bool   `xml:"enabled"`
	Flow    string `xml:"flow"`
}

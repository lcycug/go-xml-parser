package models

// ClassAccesses is used to tell if a Apex Class is visible for some Profile
type ClassAccesses struct {
	ApexClass string `xml:"apexClass"`
	Enabled   bool   `xml:"enabled"`
}

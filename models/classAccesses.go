package models

import "strings"

// ClassAccesses is used to tell if a Apex Class is visible for some Profile
type ClassAccesses struct {
	ApexClass string `xml:"apexClass"`
	Enabled   bool   `xml:"enabled"`
}

type Classes []*ClassAccesses

// Implement Sort interface
func (c Classes) Len() int      { return len(c) }
func (c Classes) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// ByAppName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Classes value.
type ByClassName struct{ Classes }

func (n ByClassName) Less(i, j int) bool {
	return strings.Compare(n.Classes[i].ApexClass, n.Classes[j].ApexClass) < 0
}

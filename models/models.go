package models

import "strings"

type Profile struct {
	Apps        []ApplicationVisibilities `xml:"applicationVisibilities"`
	Layouts     []LayoutAssignments       `xml:"layoutAssignments"`
	Tabs        []TabVisibilities         `xml:"tabVisibilities"`
	Permissions []UserPermissions         `xml:"userPermissions"`
	License     UserLicense               `xml:"userLicense"`
}

func (p Profile) Len() int { return len(p.Apps) }
func (p Profile) Less(i, j int) bool {
	return strings.Compare(p.Apps[i].Application, p.Apps[j].Application) < 0
}
func (p Profile) Swap(i, j int) {
	p.Apps[i].Application, p.Apps[j].Application = p.Apps[j].Application, p.Apps[i].Application
}

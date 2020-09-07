package models

import "strings"

type Profile struct {
	Apps       []ApplicationVisibilities `xml:"applicationVisibilities,omitempty"`
	Layouts    []LayoutAssignments       `xml:"layoutAssignments,omitempty"`
	Tabs       []TabVisibilities         `xml:"tabVisibilities,omitempty"`
	UserPerms  []UserPermissions         `xml:"userPermissions,omitempty"`
	License    UserLicense               `xml:"userLicense,omitempty"`
	FieldPerms []FieldPermissions        `xml:"fieldPermissions,omitempty"`
}

func (p Profile) Len() int { return len(p.Apps) }
func (p Profile) Less(i, j int) bool {
	return strings.Compare(p.Apps[i].Application, p.Apps[j].Application) < 0
}
func (p Profile) Swap(i, j int) {
	p.Apps[i].Application, p.Apps[j].Application = p.Apps[j].Application, p.Apps[i].Application
}

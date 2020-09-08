package models

import "strings"

// Profile is a collection of all settings for this Profile.
type Profile struct {
	Apps        []ApplicationVisibilities `xml:"applicationVisibilities,omitempty"`
	Classes     []ClassAccesses           `xml:"classAccesses,omitempty"`
	Custom      Custom                    `xml:"custom,omitempty"`
	FieldPerms  []FieldPermissions        `xml:"fieldPermissions,omitempty"`
	Flows       []FlowAccesses            `xml:"flowAccesses,omitempty"`
	Layouts     []LayoutAssignments       `xml:"layoutAssignments,omitempty"`
	License     UserLicense               `xml:"userLicense,omitempty"`
	IPs         LoginIPRanges             `xml:"loginIpRange,omitempty"`
	ObjectPerms []ObjectPermissions       `xml:"objectPermissions,omitempty"`
	Pages       []PageAccesses            `xml:"pageAccesses,omitempty"`
	RecordTypes []RecordTypeVisibilities  `xml:"recordTypeVisibilities,omitempty"`
	Tabs        []TabVisibilities         `xml:"tabVisibilities,omitempty"`
	UserPerms   []UserPermissions         `xml:"userPermissions,omitempty"`
}

func (p Profile) Len() int { return len(p.Apps) }
func (p Profile) Less(i, j int) bool {
	return strings.Compare(p.Apps[i].Application, p.Apps[j].Application) < 0
}
func (p Profile) Swap(i, j int) {
	p.Apps[i].Application, p.Apps[j].Application = p.Apps[j].Application, p.Apps[i].Application
}

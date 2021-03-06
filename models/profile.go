package models

import "github.com/fatih/structs"

type Profile interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
	First() interface{}
	Name() string
	// Fields() []*structs.Field
	Values() map[*structs.Field][]interface{}
}

// ApplicationVisibilities is used to tell if an application is visible and
// defaulted to a Profile.
type ApplicationVisibilities struct {
	Application string `xml:"application"`
	Default     bool   `xml:"default"`
	Visible     bool   `xml:"visible"`
}

// ClassAccesses is used to tell if a Apex Class is visible for some Profile
type ClassAccesses struct {
	ApexClass string `xml:"apexClass"`
	Enabled   bool   `xml:"enabled"`
}

// Custom is used to tell if a Profile is a customized one.
type Custom bool

// FieldPermissions is used to store field permissions for this Profile.
type FieldPermissions struct {
	Field    string `xml:"field"`
	Editable bool   `xml:"editable"`
	Readable bool   `xml:"readable"`
}

// FlowAccesses is used to tell if a flow is visible to a Profile
type FlowAccesses struct {
	Flow    string `xml:"flow"`
	Enabled bool   `xml:"enabled"`
}

// LayoutAssignments is used to tell if a Layout is specified for a Profile
type LayoutAssignments struct {
	Layout string `xml:"layout"`
}

// LoginIPRanges is used to specify an available IP range for a specific
// Profile.
type LoginIPRanges struct {
	EndAddress   string `xml:"endAddress"`
	StartAddress string `xml:"startAddress"`
}

// ObjectPermissions is used to store a list of permissions for a SObject.
type ObjectPermissions struct {
	Object           string `xml:"object"`
	AllowCreate      bool   `xml:"allowCreate"`
	AllowDelete      bool   `xml:"allowDelete"`
	AllowEdit        bool   `xml:"allowEdit"`
	AllowRead        bool   `xml:"allowRead"`
	ModifyAllRecords bool   `xml:"modifyAllRecords"`
	ViewAllRecords   bool   `xml:"viewAllRecords"`
}

// PageAccesses is used to specify the accessibility of a Visualforce Page.
type PageAccesses struct {
	ApexPage string `xml:"apexPage"`
	Enabled  bool   `xml:"enabled"`
}

// RecordTypeVisibilities is used to specify the properties of a Record Type.
type RecordTypeVisibilities struct {
	RecordType string `xml:"recordType"`
	Default    bool   `xml:"default"`
	Visible    bool   `xml:"visible"`
}

// TabVisibilities is used to store visibilities for Tabs
type TabVisibilities struct {
	Tab        string `xml:"tab"`
	Visibility string `xml:"visibility"`
}

// UserLicense is used to tell what user license is used for this Profile
type UserLicense string

// UserPermissions is used to store User Permissions specified for the Profile.
type UserPermissions struct {
	Name    string `xml:"name"`
	Enabled bool   `xml:"enabled"`
}

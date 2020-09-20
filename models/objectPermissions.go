package models

// ObjectPermissions is used to store a list of permissions for a SObject.
type ObjectPermissions struct {
	Name             string `xml:"object"`
	AllowCreate      bool   `xml:"allowCreate"`
	AllowDelete      bool   `xml:"allowDelete"`
	AllowEdit        bool   `xml:"allowEdit"`
	AllowRead        bool   `xml:"allowRead"`
	ModifyAllRecords bool   `xml:"modifyAllRecords"`
	ViewAllRecords   bool   `xml:"viewAllRecords"`
}

package models

// ObjectPermissions is used to store a list of permissions for a SObject.
type ObjectPermissions struct {
	AllowCreate      bool   `xml:"allowCreate"`
	AllowDelete      bool   `xml:"allowDelete"`
	AllowEdit        bool   `xml:"allowEdit"`
	AllowRead        bool   `xml:"allowRead"`
	ModifyAllRecords bool   `xml:"modifyAllRecords"`
	ViewAllRecords   bool   `xml:"viewAllRecords"`
	Object           string `xml:"object"`
}

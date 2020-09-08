package models

// FieldPermissions is used to store field permissions for this Profile.
type FieldPermissions struct {
	Editable bool   `xml:"editable"`
	Field    string `xml:"field"`
	Readable bool   `xml:"readable"`
}

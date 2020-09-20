package models

// FieldPermissions is used to store field permissions for this Profile.
type FieldPermissions struct {
	Name     string `xml:"field"`
	Editable bool   `xml:"editable"`
	Readable bool   `xml:"readable"`
}

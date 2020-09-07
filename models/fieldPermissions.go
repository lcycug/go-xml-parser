package models

type FieldPermissions struct {
	Editable bool   `xml:"editable"`
	Field    string `xml:"field"`
	Readable bool   `xml:"readable"`
}

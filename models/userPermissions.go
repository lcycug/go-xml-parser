package models

//UserPermissions is used to store User Permissions specified for the Profile.
type UserPermissions struct {
	Name    string `xml:"name"`
	Enabled bool   `xml:"enabled"`
}

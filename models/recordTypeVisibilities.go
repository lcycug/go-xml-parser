package models

// RecordTypeVisibilities is used to specify the properties of a Record Type.
type RecordTypeVisibilities struct {
	Name    string `xml:"recordType"`
	Default bool   `xml:"default"`
	Visible bool   `xml:"visible"`
}

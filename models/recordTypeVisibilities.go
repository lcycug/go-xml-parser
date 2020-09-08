package models

// RecordTypeVisibilities is used to specify the properties of a Record Type.
type RecordTypeVisibilities struct {
	Default    bool   `xml:"default"`
	Visible    bool   `xml:"visible"`
	RecordType string `xml:"recordType"`
}

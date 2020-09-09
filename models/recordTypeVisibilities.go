package models

import "strings"

// RecordTypeVisibilities is used to specify the properties of a Record Type.
type RecordTypeVisibilities struct {
	Default    bool   `xml:"default"`
	Visible    bool   `xml:"visible"`
	RecordType string `xml:"recordType"`
}

type RecordTypes []*RecordTypeVisibilities

// Implement Sort interface
func (p RecordTypes) Len() int      { return len(p) }
func (p RecordTypes) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// ByRecordTypeName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded RecordTypes value.
type ByRecordTypeName struct{ RecordTypes }

func (n ByRecordTypeName) Less(i, j int) bool {
	return strings.Compare(n.RecordTypes[i].RecordType, n.RecordTypes[j].RecordType) < 0
}

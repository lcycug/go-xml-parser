package recordTypes

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS       string                           `xml:"xmlns,attr,omitempty"`
	RecordTypes []*models.RecordTypeVisibilities `xml:"recordTypeVisibilities,omitempty"`
}

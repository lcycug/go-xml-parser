package recordTypes

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	RecordTypes []*models.RecordTypeVisibilities `xml:"recordTypeVisibilities,omitempty"`
}

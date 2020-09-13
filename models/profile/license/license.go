package license

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS       string              `xml:"xmlns,attr,omitempty"`
	UserLicense *models.UserLicense `xml:"userLicense,omitempty"`
}

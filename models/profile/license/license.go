package license

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	UserLicense *models.UserLicense `xml:"userLicense,omitempty"`
}

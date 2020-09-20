package ips

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	XMLNS string                  `xml:"xmlns,attr,omitempty"`
	IPs   []*models.LoginIPRanges `xml:"loginIpRange,omitempty"`
}

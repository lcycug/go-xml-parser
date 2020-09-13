package ips

import "github.com/lcycug/go-xml-parser/models"

type Profile struct {
	IPs []*models.LoginIPRanges `xml:"loginIpRange,omitempty"`
}

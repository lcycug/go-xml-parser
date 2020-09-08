package models

// LoginIpRanges is used to specify an available IP range for a specific Profile.
type LoginIpRanges struct {
	EndAddress   string `xml:"endAddress"`
	StartAddress string `xml:"startAddress"`
}

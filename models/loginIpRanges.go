package models

// LoginIPRanges is used to specify an available IP range for a specific Profile.
type LoginIPRanges struct {
	EndAddress   string `xml:"endAddress"`
	StartAddress string `xml:"startAddress"`
}

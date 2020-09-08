package models

import "fmt"

// UserLicense is used to tell what user license is used for this Profile
type UserLicense string

// This method is used to return a formatted string.
func (l UserLicense) String() string {
	return fmt.Sprintf("UserLicense: %s", string(l))
}

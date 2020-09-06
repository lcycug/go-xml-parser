package models

import "fmt"

type UserLicense string

func (l UserLicense) String() string {
	return fmt.Sprintf("UserLicense: %s", string(l))
}

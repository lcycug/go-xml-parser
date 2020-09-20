package utils

// NewNextAlphabetInstance is used to instance an alphabet generator.
func NewNextAlphabetInstance() func() string {
	var i int32 = 64
	return func() string {
		i++
		return string(i)
	}
}

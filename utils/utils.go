package utils

import (
	"log"
)

//Errors is a customized error struct type
type Errors struct {
	Type  int
	Error error
}

const (
	ERROR = iota
	WARN
	SUCCESS
)

const (
	SPACE = " "
	DOT   = "."
)

func LogFatal(cErr string, err error) {
	if err != nil {
		log.Fatal(cErr, err)
	}
}

// InstanceNextAlphabet is used to instance a alphabet generator.
func InstanceNextAlphabet() func() string {
	var i int32 = 64
	return func() string {
		i++
		return string(i)
	}
}

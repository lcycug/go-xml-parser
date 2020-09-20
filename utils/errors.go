package utils

import "log"

// Errors is a customized error struct type
type Errors struct {
	Type  int
	Error error
}

func LogFatal(cErr string, err error) {
	if err != nil {
		log.Fatal(cErr, err)
	}
}

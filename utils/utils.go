package utils

import (
	"log"
)

func LogFatal(cErr string, err error) {
	if err != nil {
		log.Fatal(cErr, err)
	}
}

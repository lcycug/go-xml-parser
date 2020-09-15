package utils

import (
	"log"
	"reflect"
	"strings"

	"github.com/fatih/structs"
)

func LogFatal(cErr string, err error) {
	if err != nil {
		log.Fatal(cErr, err)
	}
}

func GetFileName(c interface{}) string {
	var n string
	if structs.IsStruct(c) {
		n = structs.Name(c)
	} else {
		n = reflect.TypeOf(c).String()
	}
	ns := strings.Split(n, ".")
	if len(ns) > 1 {
		n = ns[1]
	} else {
		n = ns[0]
	}
	return strings.ToLower(string(n[0])) + n[1:] + ".xml"
}

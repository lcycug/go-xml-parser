package main

import (
	"encoding/xml"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lcycug/go-xml-parser/models"
	"io/ioutil"
	"os"
	"sort"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	resp, err := ioutil.ReadFile(os.Getenv("ADMIN_PROFILE_PATH"))
	if err != nil {
		panic(err)
	}
	var admin models.Profile
	err = xml.Unmarshal(resp, &admin)
	if err != nil {
		panic(err)
	}
	sort.Sort(admin)
	fmt.Println(admin)
}

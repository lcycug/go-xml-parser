package main

import (
	"encoding/xml"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/joho/godotenv"
	"github.com/lcycug/go-xml-parser/models"
	"github.com/lcycug/go-xml-parser/models/profile/apps"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	var fs []*excelize.File
	files, err := ioutil.ReadDir(os.Getenv("PROFILE_PATH"))
	if err != nil {
		log.Fatal("Failed to read directory:", err)
	}
	for _, fileInfo := range files {
		var p apps.Profile
		xls := excelize.NewFile()
		fs = append(fs, xls)
		ss := strings.Split(fileInfo.Name(), ".")
		fmt.Println("You are manipulating file:", fileInfo.Name())
		if len(ss) != 3 {
			continue
		}
		f, err := ioutil.ReadFile(os.Getenv("PROFILE_PATH") + fileInfo.Name())
		if err != nil {
			log.Fatal("Failed to read file:", err)
		}
		err = xml.Unmarshal(f, &p)
		if err != nil {
			log.Fatal("Failed to unmarshal xml:", err)
		}
		sort.Sort(models.ByAppName{Apps: p.Apps})

		err = os.Mkdir(filepath.Join([]string{os.Getenv("PROFILE_PATH"), ss[0]}...), os.ModePerm)
		if err != nil {
			log.Fatal("Failed to create a new directory: ", err)
		}

		nf, _ := xml.MarshalIndent(p, "", "  ")
		_ = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"), ss[0], "Apps.xml"), nf, 0644)
		break
	}
}

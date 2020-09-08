package main

import (
	"encoding/xml"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/joho/godotenv"
	"github.com/lcycug/go-xml-parser/models"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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
		var p models.Profile
		xls := excelize.NewFile()
		fs = append(fs, xls)
		ss := strings.Split(fileInfo.Name(), ".")
		if len(ss) != 3 {
			return
		}
		f, err := ioutil.ReadFile(os.Getenv("PROFILE_PATH") + fileInfo.Name())
		if err != nil {
			log.Fatal("Failed to read file::", err)
		}
		err = xml.Unmarshal(f, &p)
		if err != nil {
			log.Fatal("Failed to unmarshal xml:", err)
		}
		if len(p.Apps) != 0 {
			xls.NewSheet("Application")
			xls.SetCellStr("Application", "A1", "Application")
			xls.SetCellStr("Application", "B1", "Default")
			xls.SetCellStr("Application", "C1", "Visible")
		}
		for i, app := range p.Apps {
			xls.SetCellStr("Application", "A"+strconv.Itoa(i+2), app.Application)
			xls.SetCellBool("Application", "B"+strconv.Itoa(i+2), app.Default)
			xls.SetCellBool("Application", "C"+strconv.Itoa(i+2), app.Visible)
		}
		xls.DeleteSheet("Sheet1")
		if err = xls.SaveAs(ss[0] + ".xlsx"); err != nil {
			log.Fatal("Failed to save .xlsx file", err)
		}
		break
	}
}

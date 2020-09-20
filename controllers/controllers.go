package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/lcycug/go-xml-parser/models/profile/apps"
	"github.com/lcycug/go-xml-parser/models/profile/classes"
	"github.com/lcycug/go-xml-parser/models/profile/custom"
	"github.com/lcycug/go-xml-parser/models/profile/fieldPerms"
	"github.com/lcycug/go-xml-parser/models/profile/flows"
	"github.com/lcycug/go-xml-parser/models/profile/ips"
	"github.com/lcycug/go-xml-parser/models/profile/layouts"
	"github.com/lcycug/go-xml-parser/models/profile/license"
	"github.com/lcycug/go-xml-parser/models/profile/objectPerms"
	"github.com/lcycug/go-xml-parser/models/profile/pages"
	"github.com/lcycug/go-xml-parser/models/profile/recordTypes"
	"github.com/lcycug/go-xml-parser/models/profile/tabs"
	"github.com/lcycug/go-xml-parser/models/profile/userPerms"
	"github.com/lcycug/go-xml-parser/utils"
)

// SplitProfile aims to split a Profile into pieces of components.
func SplitProfile(path string, fi os.FileInfo) utils.Errors {
	var (
		ap  apps.Profile
		clp classes.Profile
		cup custom.Profile
		fip fieldPerms.Profile
		flp flows.Profile
		lp  layouts.Profile
		ip  ips.Profile
		op  objectPerms.Profile
		pp  pages.Profile
		rp  recordTypes.Profile
		tp  tabs.Profile
		ulp license.Profile
		upp userPerms.Profile
	)

	fmt.Println("You are manipulating file:", fi.Name())

	xls := excelize.NewFile()
	ss := strings.Split(fi.Name(), utils.DOT)

	if len(ss) != 3 {
		return utils.Errors{
			Type:  utils.WARN,
			Error: fmt.Errorf("invalid file: %s", fi.Name()),
		}
	}

	f, err := ioutil.ReadFile(path + fi.Name())
	utils.LogFatal("Failed to read file:", err)

	err = os.RemoveAll(filepath.Join([]string{path, ss[0]}...))
	utils.LogFatal("Failed to remove folder", err)

	err = os.Mkdir(filepath.Join([]string{path, ss[0]}...), os.ModePerm)
	utils.LogFatal("Failed to create a new directory: ", err)

	CreateXML(path, ss, f, xls, &ap)
	CreateXML(path, ss, f, xls, &clp)
	CreateXML(path, ss, f, xls, &cup)
	CreateXML(path, ss, f, xls, &fip)
	CreateXML(path, ss, f, xls, &flp)
	CreateXML(path, ss, f, xls, &lp)
	CreateXML(path, ss, f, xls, &ip)
	CreateXML(path, ss, f, xls, &op)
	CreateXML(path, ss, f, xls, &pp)
	CreateXML(path, ss, f, xls, &rp)
	CreateXML(path, ss, f, xls, &tp)
	CreateXML(path, ss, f, xls, &ulp)
	CreateXML(path, ss, f, xls, &upp)

	xls.DeleteSheet("Sheet1")
	_ = xls.SaveAs(filepath.Join(path, ss[0], ss[0]+".xlsx"))
	return utils.Errors{}
}

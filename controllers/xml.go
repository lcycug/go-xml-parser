package controllers

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/fatih/structs"
	"github.com/lcycug/go-xml-parser/models"
	"github.com/lcycug/go-xml-parser/models/apps"
	"github.com/lcycug/go-xml-parser/models/classes"
	"github.com/lcycug/go-xml-parser/models/fieldPerms"
	"github.com/lcycug/go-xml-parser/models/flows"
	"github.com/lcycug/go-xml-parser/models/layouts"
	"github.com/lcycug/go-xml-parser/models/objectPerms"
	"github.com/lcycug/go-xml-parser/models/pages"
	"github.com/lcycug/go-xml-parser/models/recordTypes"
	"github.com/lcycug/go-xml-parser/models/tabs"
	"github.com/lcycug/go-xml-parser/models/userPerms"
	"github.com/lcycug/go-xml-parser/utils"
)

// CreateXML is used to create a xml file for a type of Profile component.
// ss: Split Profile Name parts, [Admin, profile-meta, xml]
// f: File
// v: One type of Profile component
func CreateXML(path string, ss []string, f []byte, xls *excelize.File, pn int) {

	var p models.Profile

	switch pn {
	case utils.AppProfile:
		p = new(apps.Profile)
		break
	case utils.ClassProfile:
		p = new(classes.Profile)
		break
	case utils.FieldPermProfile:
		p = new(fieldPerms.Profile)
		break
	case utils.FlowProfile:
		p = new(flows.Profile)
		break
	case utils.LayoutProfile:
		p = new(layouts.Profile)
		break
	case utils.ObjectPermProfile:
		p = new(objectPerms.Profile)
		break
	case utils.PageProfile:
		p = new(pages.Profile)
		break
	case utils.RecordTypeProfile:
		p = new(recordTypes.Profile)
		break
	case utils.TabProfile:
		p = new(tabs.Profile)
		break
	case utils.UserPermProfile:
		p = new(userPerms.Profile)
		break
	default:
		return
	}

	utils.LogFatal("Failed to marshal xml:", xml.Unmarshal(f, &p))
	if p.Len() == 0 {
		return
	}
	sort.Sort(p)
	n := getXMLFileName(p.First(), ss[1:])

	createSheet(p, xls)

	nf, err := xml.MarshalIndent(p, "", utils.SPACE+utils.SPACE+utils.
		SPACE+utils.SPACE)
	utils.LogFatal("Failed to marshal xml:", err)

	nf = []byte(xml.Header + string(nf))
	err = ioutil.WriteFile(filepath.Join(path, ss[0], n), nf, 0644)
	utils.LogFatal("Failed to write xml:", err)
}

// getXMLFileName is a private method to retrieve a Name from a struct or a type.
// c is struct or a type
// ss is a slice of string
func getXMLFileName(c interface{}, ss []string) string {
	var n string
	if structs.IsStruct(c) {
		n = structs.Name(c)
	} else {
		n = reflect.TypeOf(c).String()
	}
	ns := strings.Split(n, utils.DOT)
	if len(ns) > 1 {
		n = ns[1]
	} else {
		n = ns[0]
	}
	return strings.ToLower(string(n[0])) + n[1:] + utils.DOT + strings.Join(ss,
		utils.DOT)
}

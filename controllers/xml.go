package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/fatih/structs"
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
func CreateXML(path string, ss []string, f []byte, xls *excelize.File,
	v interface{}) {

	var sd, n interface{}

	err := xml.Unmarshal(f, &v)
	utils.LogFatal("Failed to unmarshal xml:", err)

	switch v.(type) {
	case *apps.Profile:
		d := v.(*apps.Profile)
		if len(d.Apps) == 0 {
			return
		}
		sort.Sort(apps.ByName{Profile: *d})
		n = d.Apps[0]
		sd = d
		break
	case *classes.Profile:
		d := v.(*classes.Profile)
		if len(d.Classes) == 0 {
			return
		}
		sort.Sort(classes.ByName{Profile: *d})
		n = d.Classes[0]
		sd = d
		break
	case *fieldPerms.Profile:
		d := v.(*fieldPerms.Profile)
		if len(d.FieldPerms) == 0 {
			return
		}
		sort.Sort(fieldPerms.ByName{Profile: *d})
		n = d.FieldPerms[0]
		sd = d
		break
	case *flows.Profile:
		d := v.(*flows.Profile)
		if len(d.Flows) == 0 {
			return
		}
		sort.Sort(flows.ByName{Profile: *d})
		n = d.Flows[0]
		sd = d
		break
	case *layouts.Profile:
		d := v.(*layouts.Profile)
		if len(d.Layouts) == 0 {
			return
		}
		sort.Sort(layouts.ByName{Profile: *d})
		n = d.Layouts[0]
		sd = d
		break
	case *objectPerms.Profile:
		d := v.(*objectPerms.Profile)
		if len(d.ObjectPerms) == 0 {
			return
		}
		sort.Sort(objectPerms.ByName{Profile: *d})
		n = d.ObjectPerms[0]
		sd = d
		break
	case *pages.Profile:
		d := v.(*pages.Profile)
		if len(d.Pages) == 0 {
			return
		}
		sort.Sort(pages.ByName{Profile: *d})
		n = d.Pages[0]
		sd = d
		break
	case *recordTypes.Profile:
		d := v.(*recordTypes.Profile)
		if len(d.RecordTypes) == 0 {
			return
		}
		sort.Sort(recordTypes.ByName{Profile: *d})
		n = d.RecordTypes[0]
		sd = d
		break
	case *tabs.Profile:
		d := v.(*tabs.Profile)
		if len(d.Tabs) == 0 {
			return
		}
		sort.Sort(tabs.ByName{Profile: *d})
		n = d.Tabs[0]
		sd = d
		break
	case *userPerms.Profile:
		d := v.(*userPerms.Profile)
		if len(d.UserPerms) == 0 {
			return
		}
		sort.Sort(userPerms.ByName{Profile: *d})
		n = d.UserPerms[0]
		sd = d
		break
	default:
		fmt.Println("No matched type found!")
		return
	}

	createSheet(sd, xls)

	nf, err := xml.MarshalIndent(sd, "", utils.SPACE+utils.SPACE+utils.
		SPACE+utils.SPACE)
	utils.LogFatal("Failed to marshal xml:", err)

	nf = []byte(xml.Header + string(nf))
	err = ioutil.WriteFile(filepath.Join(path, ss[0], getXMLFileName(n,
		ss[1:])), nf, 0644)
	utils.LogFatal("Failed to write xml:", err)
}

// getFileName is a private method to retrieve a Name from a struct or a type.
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

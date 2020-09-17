package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/fatih/structs"
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

//Errors is a customized error struct type
type Errors struct {
	Type  int
	Error error
}

const (
	ERROR = iota - 2
	WARN
	SUCCESS
)

const (
	space = " "
	dot   = "."
)

//CreateFile is used to create a xml file for a type of Profile component.
//ss: Split Profile Name parts, [Admin, profile-meta, xml]
//f: File
//v: One type of Profile component
func CreateFile(ss []string, f []byte, v interface{}) {
	var sd interface{}
	var n interface{}
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
		sd = interface{}(d)
		break
	case *classes.Profile:
		d := v.(*classes.Profile)
		if len(d.Classes) == 0 {
			return
		}
		sort.Sort(classes.ByName{Profile: *d})
		n = d.Classes[0]
		sd = interface{}(d)
		break
	case *fieldPerms.Profile:
		d := v.(*fieldPerms.Profile)
		if len(d.FieldPerms) == 0 {
			return
		}
		sort.Sort(fieldPerms.ByName{Profile: *d})
		n = d.FieldPerms[0]
		sd = interface{}(d)
		break
	case *flows.Profile:
		d := v.(*flows.Profile)
		if len(d.Flows) == 0 {
			return
		}
		sort.Sort(flows.ByName{Profile: *d})
		n = d.Flows[0]
		sd = interface{}(d)
		break
	case *layouts.Profile:
		d := v.(*layouts.Profile)
		if len(d.Layouts) == 0 {
			return
		}
		sort.Sort(layouts.ByName{Profile: *d})
		n = d.Layouts[0]
		sd = interface{}(d)
		break
	case *objectPerms.Profile:
		d := v.(*objectPerms.Profile)
		if len(d.ObjectPerms) == 0 {
			return
		}
		sort.Sort(objectPerms.ByName{Profile: *d})
		n = d.ObjectPerms[0]
		sd = interface{}(d)
		break
	case *pages.Profile:
		d := v.(*pages.Profile)
		if len(d.Pages) == 0 {
			return
		}
		sort.Sort(pages.ByName{Profile: *d})
		n = d.Pages[0]
		sd = interface{}(d)
		break
	case *recordTypes.Profile:
		d := v.(*recordTypes.Profile)
		if len(d.RecordTypes) == 0 {
			return
		}
		sort.Sort(recordTypes.ByName{Profile: *d})
		n = d.RecordTypes[0]
		sd = interface{}(d)
		break
	case *tabs.Profile:
		d := v.(*tabs.Profile)
		if len(d.Tabs) == 0 {
			return
		}
		sort.Sort(tabs.ByName{Profile: *d})
		n = d.Tabs[0]
		sd = interface{}(d)
		break
	case *userPerms.Profile:
		d := v.(*userPerms.Profile)
		if len(d.UserPerms) == 0 {
			return
		}
		sort.Sort(userPerms.ByName{Profile: *d})
		n = d.UserPerms[0]
		sd = interface{}(d)
		break
	default:
		fmt.Println("No matched type found!")
		return
	}

	nf, err := xml.MarshalIndent(sd, "", space+space+space+space)
	utils.LogFatal("Failed to marshal xml:", err)

	nf = []byte(xml.Header + string(nf))
	err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
		ss[0], getFileName(n, ss[1:])), nf, 0644)
	utils.LogFatal("Failed to write xml:", err)
}

//SplitProfile aims to split a Profile into pieces according to different
//attributes.
func SplitProfile(fi os.FileInfo) Errors {
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
	ss := strings.Split(fi.Name(), dot)
	fmt.Println("You are manipulating file:", fi.Name())
	if len(ss) != 3 {
		return Errors{
			Type:  WARN,
			Error: fmt.Errorf("invalid file: %s", fi.Name()),
		}
	}
	f, err := ioutil.ReadFile(os.Getenv("PROFILE_PATH") + fi.Name())
	utils.LogFatal("Failed to read file:", err)

	err = os.RemoveAll(filepath.Join([]string{os.Getenv("PROFILE_PATH"),
		ss[0]}...))
	utils.LogFatal("Failed to remove folder", err)
	err = os.Mkdir(filepath.Join([]string{os.Getenv("PROFILE_PATH"),
		ss[0]}...), os.ModePerm)
	utils.LogFatal("Failed to create a new directory: ", err)

	CreateFile(ss, f, &ap)
	CreateFile(ss, f, &clp)
	CreateFile(ss, f, &cup)
	CreateFile(ss, f, &fip)
	CreateFile(ss, f, &flp)
	CreateFile(ss, f, &lp)
	CreateFile(ss, f, &ip)
	CreateFile(ss, f, &op)
	CreateFile(ss, f, &pp)
	CreateFile(ss, f, &rp)
	CreateFile(ss, f, &tp)
	CreateFile(ss, f, &ulp)
	CreateFile(ss, f, &upp)
	return Errors{}
}

//getFileName is a private method to retrieve a Name from a struct or a type.
func getFileName(c interface{}, ss []string) string {
	var n string
	if structs.IsStruct(c) {
		n = structs.Name(c)
	} else {
		n = reflect.TypeOf(c).String()
	}
	ns := strings.Split(n, dot)
	if len(ns) > 1 {
		n = ns[1]
	} else {
		n = ns[0]
	}
	return strings.ToLower(string(n[0])) + n[1:] + dot + strings.Join(ss, dot)
}

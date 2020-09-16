package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

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

//CreateFile is used to create a xml file for a type of Profile component.
//pn: Profile Name
//f: File
//v: One type of Profile component
func CreateFile(pn string, f []byte, v interface{}) {
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

	nf, err := xml.MarshalIndent(sd, "", "  ")
	utils.LogFatal("Failed to marshal xml:", err)

	nf = []byte(xml.Header + string(nf))
	err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
		pn, utils.GetFileName(n)), nf, 0644)
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
	ss := strings.Split(fi.Name(), ".")
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

	CreateFile(ss[0], f, &ap)
	CreateFile(ss[0], f, &clp)
	CreateFile(ss[0], f, &cup)
	CreateFile(ss[0], f, &fip)
	CreateFile(ss[0], f, &flp)
	CreateFile(ss[0], f, &lp)
	CreateFile(ss[0], f, &ip)
	CreateFile(ss[0], f, &op)
	CreateFile(ss[0], f, &pp)
	CreateFile(ss[0], f, &rp)
	CreateFile(ss[0], f, &tp)
	CreateFile(ss[0], f, &ulp)
	CreateFile(ss[0], f, &upp)
	return Errors{}
}

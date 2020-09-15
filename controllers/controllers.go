package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/lcycug/go-xml-parser/models/profile/apps"
	"github.com/lcycug/go-xml-parser/models/profile/classes"
	"github.com/lcycug/go-xml-parser/models/profile/fieldPerms"
	"github.com/lcycug/go-xml-parser/models/profile/flows"
	"github.com/lcycug/go-xml-parser/models/profile/layouts"
	"github.com/lcycug/go-xml-parser/models/profile/objectPerms"
	"github.com/lcycug/go-xml-parser/models/profile/pages"
	"github.com/lcycug/go-xml-parser/models/profile/recordTypes"
	"github.com/lcycug/go-xml-parser/models/profile/tabs"
	"github.com/lcycug/go-xml-parser/models/profile/userPerms"
	"github.com/lcycug/go-xml-parser/utils"
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

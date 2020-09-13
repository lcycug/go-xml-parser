package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/fatih/structs"
	"github.com/joho/godotenv"
	"github.com/lcycug/go-xml-parser/models"
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
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	fileInfos, err := ioutil.ReadDir(os.Getenv("PROFILE_PATH"))
	logFatal("Failed to read directory:", err)
	for _, fi := range fileInfos {
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
			continue
		}
		f, err := ioutil.ReadFile(os.Getenv("PROFILE_PATH") + fi.Name())
		logFatal("Failed to read file:", err)

		err = os.Mkdir(filepath.Join([]string{os.Getenv("PROFILE_PATH"),
			ss[0]}...), os.ModePerm)
		logFatal("Failed to create a new directory: ", err)

		err = xml.Unmarshal(f, &ap)
		logFatal("Failed to unmarshal xml:", err)
		if len(f) > 0 {
			sort.Sort(models.ByAppName{Apps: ap.Apps})
			nf, err := xml.MarshalIndent(ap, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(ap.Apps[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &clp)
		logFatal("Failed to unmarshal xml:", err)
		if len(clp.Classes) > 0 {
			sort.Sort(models.ByClassName{Classes: clp.Classes})
			nf, err := xml.MarshalIndent(clp, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(clp.Classes[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &cup)
		logFatal("Failed to unmarshal xml:", err)
		if cup.Custom != nil {
			nf, err := xml.MarshalIndent(cup, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(cup.Custom)), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &fip)
		logFatal("Failed to unmarshal xml:", err)
		if len(fip.FieldPerms) > 0 {
			sort.Sort(models.ByFieldName{FieldPerms: fip.FieldPerms})
			nf, err := xml.MarshalIndent(fip, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(fip.FieldPerms[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &flp)
		logFatal("Failed to unmarshal xml:", err)
		if len(flp.Flows) > 0 {
			sort.Sort(models.ByFlowName{Flows: flp.Flows})
			nf, err := xml.MarshalIndent(flp, "", "  ")
			logFatal("Failed to marshal xml:", err)
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(flp.Flows[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &lp)
		logFatal("Failed to unmarshal xml:", err)
		if len(lp.Layouts) > 0 {
			sort.Sort(models.ByLayoutName{Layouts: lp.Layouts})
			nf, err := xml.MarshalIndent(lp, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(lp.Layouts[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &ip)
		logFatal("Failed to unmarshal xml:", err)
		if len(ip.IPs) > 0 {
			nf, err := xml.MarshalIndent(ip, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(ip.IPs[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &op)
		logFatal("Failed to unmarshal xml:", err)
		if len(op.ObjectPerms) > 0 {
			sort.Sort(models.ByObjectName{ObjectPerms: op.ObjectPerms})
			nf, err := xml.MarshalIndent(op, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(op.ObjectPerms[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &pp)
		logFatal("Failed to unmarshal xml:", err)
		if len(pp.Pages) > 0 {
			sort.Sort(models.ByPageName{Pages: pp.Pages})
			nf, err := xml.MarshalIndent(pp, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(pp.Pages[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &rp)
		logFatal("Failed to unmarshal xml:", err)
		if len(rp.RecordTypes) > 0 {
			sort.Sort(models.ByRecordTypeName{RecordTypes: rp.RecordTypes})
			nf, err := xml.MarshalIndent(rp, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(rp.RecordTypes[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &tp)
		logFatal("Failed to unmarshal xml:", err)
		if len(tp.Tabs) > 0 {
			sort.Sort(models.ByTabName{Tabs: tp.Tabs})
			nf, err := xml.MarshalIndent(tp, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(tp.Tabs[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &ulp)
		logFatal("Failed to unmarshal xml:", err)
		if ulp.UserLicense != nil {
			nf, err := xml.MarshalIndent(ulp, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(ulp.UserLicense)), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		err = xml.Unmarshal(f, &upp)
		logFatal("Failed to unmarshal xml:", err)
		if len(upp.UserPerms) > 0 {
			sort.Sort(models.ByUserPermName{UserPerms: upp.UserPerms})
			nf, err := xml.MarshalIndent(upp, "", "  ")
			logFatal("Failed to marshal xml:", err)
			nf = []byte(xml.Header + string(nf))
			err = ioutil.WriteFile(filepath.Join(os.Getenv("PROFILE_PATH"),
				ss[0], getFileName(upp.UserPerms[0])), nf, 0644)
			logFatal("Failed to write xml:", err)
		}

		break
	}
}

func logFatal(cErr string, err error) {
	if err != nil {
		log.Fatal(cErr, err)
	}
}

func getFileName(c interface{}) string {
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

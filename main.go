package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lcycug/go-xml-parser/controllers"
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

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	fileInfos, err := ioutil.ReadDir(os.Getenv("PROFILE_PATH"))
	utils.LogFatal("Failed to read directory:", err)
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
		utils.LogFatal("Failed to read file:", err)

		err = os.RemoveAll(filepath.Join([]string{os.Getenv("PROFILE_PATH"),
			ss[0]}...))
		utils.LogFatal("Failed to remove folder", err)
		err = os.Mkdir(filepath.Join([]string{os.Getenv("PROFILE_PATH"),
			ss[0]}...), os.ModePerm)
		utils.LogFatal("Failed to create a new directory: ", err)

		controllers.CreateFile(ss[0], f, &ap)
		controllers.CreateFile(ss[0], f, &clp)
		controllers.CreateFile(ss[0], f, &cup)
		controllers.CreateFile(ss[0], f, &fip)
		controllers.CreateFile(ss[0], f, &flp)
		controllers.CreateFile(ss[0], f, &lp)
		controllers.CreateFile(ss[0], f, &ip)
		controllers.CreateFile(ss[0], f, &op)
		controllers.CreateFile(ss[0], f, &pp)
		controllers.CreateFile(ss[0], f, &rp)
		controllers.CreateFile(ss[0], f, &tp)
		controllers.CreateFile(ss[0], f, &ulp)
		controllers.CreateFile(ss[0], f, &upp)
		break
	}
}

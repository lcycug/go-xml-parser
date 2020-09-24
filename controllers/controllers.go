package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/lcycug/go-xml-parser/utils"
)

// SplitProfile aims to split a Profile into pieces of components.
func SplitProfile(path string, fi os.FileInfo) utils.Errors {
	profiles := []int{
		utils.AppProfile,
		utils.ClassProfile,
		utils.CustomProfile,
		utils.FieldPermProfile,
		utils.FlowProfile,
		utils.LayoutProfile,
		utils.LicenseProfile,
		utils.IPProfile,
		utils.ObjectPermProfile,
		utils.PageProfile,
		utils.RecordTypeProfile,
		utils.TabProfile,
		utils.UserPermProfile,
	}

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

	for _, p := range profiles {
		CreateXML(path, ss, f, xls, p)
	}

	xls.DeleteSheet("Sheet1")
	_ = xls.SaveAs(filepath.Join(path, ss[0], ss[0]+".xlsx"))
	return utils.Errors{}
}

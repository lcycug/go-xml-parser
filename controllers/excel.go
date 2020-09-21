package controllers

import (
	"fmt"
	"strconv"

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

// createSheet is used to create a excel file.
// xls: pointer to a excelize.File file.
// v: data to be used to fill in the excel of a specific Profile type.
func createSheet(v interface{}, xls *excelize.File) {
	var n string
	switch v.(type) {
	case *apps.Profile:
		n = structs.Name(models.ApplicationVisibilities{})
		break
	case *fieldPerms.Profile:
		n = structs.Name(models.LayoutAssignments{})
		break
	case *flows.Profile:
		n = structs.Name(models.FlowAccesses{})
		break
	case *layouts.Profile:
		n = structs.Name(models.LayoutAssignments{})
		break
	case *objectPerms.Profile:
		n = structs.Name(models.ObjectPermissions{})
		break
	case *pages.Profile:
		n = structs.Name(models.PageAccesses{})
		break
	case *recordTypes.Profile:
		n = structs.Name(models.RecordTypeVisibilities{})
		break
	case *tabs.Profile:
		n = structs.Name(models.TabVisibilities{})
		break
	case *userPerms.Profile:
		n = structs.Name(models.UserPermissions{})
		break
	default:
		fmt.Println("No matched type found!")
		return
	}
	xls.NewSheet(n)
	setValues(xls, n, v)
}

// setValues is used to set cell values for a sheet of a workbook.
// xls is a profile specific *excelize.File file.
// sheet is for a specific type of component for a profile.
// data is used to be parsed into excel.
func setValues(xls *excelize.File, sheet string, data interface{}) {
	var err error
	var nextAlphabet func() string

	switch data.(type) {
	case *apps.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.ApplicationVisibilities{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*apps.Profile)
		for i, v := range d.Apps {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Application)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Default)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Visible)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *classes.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.ClassAccesses{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*classes.Profile)
		for i, v := range d.Classes {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.ApexClass)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Enabled)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *fieldPerms.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.FieldPermissions{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*fieldPerms.Profile)
		for i, v := range d.FieldPerms {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Field)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Editable)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Readable)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *flows.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.FlowAccesses{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*flows.Profile)
		for i, v := range d.Flows {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Flow)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Enabled)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *layouts.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.LayoutAssignments{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*layouts.Profile)
		for i, v := range d.Layouts {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Layout)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *objectPerms.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.ObjectPermissions{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*objectPerms.Profile)
		for i, v := range d.ObjectPerms {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Name)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.AllowCreate)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.AllowDelete)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.AllowEdit)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.AllowRead)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.ModifyAllRecords)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.ViewAllRecords)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *pages.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.PageAccesses{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*pages.Profile)
		for i, v := range d.Pages {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Name)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Enabled)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *recordTypes.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.RecordTypeVisibilities{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*recordTypes.Profile)
		for i, v := range d.RecordTypes {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Name)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Default)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Visible)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *tabs.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.TabVisibilities{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*tabs.Profile)
		for i, v := range d.Tabs {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Name)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Visibility)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	case *userPerms.Profile:
		nextAlphabet = utils.NewNextAlphabetInstance()
		for _, f := range structs.Fields(models.UserPermissions{}) {
			err = xls.SetCellValue(sheet, nextAlphabet()+"1", f.Name())
			utils.LogFatal("Failed to set cell value: "+f.Name(), err)
		}
		d := data.(*userPerms.Profile)
		for i, v := range d.UserPerms {
			nextAlphabet = utils.NewNextAlphabetInstance()
			err = xls.SetCellStr(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Name)
			utils.LogFatal("Failed to set cell value: ", err)
			err = xls.SetCellBool(sheet, nextAlphabet()+strconv.Itoa(i+2),
				v.Enabled)
			utils.LogFatal("Failed to set cell value: ", err)
		}
	}
}

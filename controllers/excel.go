package controllers

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/lcycug/go-xml-parser/models"
	"github.com/lcycug/go-xml-parser/utils"
)

// createSheet is used to create a excel file.
// xls: pointer to a excelize.File file.
// v: data to be used to fill in the excel of a specific Profile type.
func createSheet(p models.Profile, xls *excelize.File) {
	n := p.Name()
	xls.NewSheet(n)
	nextAlphabet := utils.NewNextAlphabetInstance()
	for f, v := range p.Values() {
		c := nextAlphabet()
		utils.LogFatal("Failed to set value:", xls.SetCellValue(n, c+"1",
			f.Name()))
		for i, vi := range v {
			utils.LogFatal("Failed to set value:", xls.SetCellValue(n,
				c+strconv.Itoa(i+2), vi))
		}
	}
}

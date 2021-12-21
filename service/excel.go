package service

import (
	"fmt"
	"resource-plan-improvement/config"
	"resource-plan-improvement/entity"
	"time"

	"github.com/xuri/excelize/v2"
)

var (
	excel = config.Conf.Excel
)

func GenExcel() (path string, err error) {
	tgtSheet := "Total"
	// open file
	dir := "files"
	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s_%s.xlsx", excel.MasterResPlanPrefix, date)
	templateFilename := fmt.Sprintf("%s.xlsx", excel.MasterResPlanPrefix)
	var file *excelize.File
	if file, err = excelize.OpenFile(dir + "/" + templateFilename); err != nil {
		return
	}
	// get data from db
	var masterResPlanItems []entity.MasterResPlanItem
	if masterResPlanItems, err = entity.FindMasterResPlan(); err != nil {
		return
	}
	// renew sheet and get writer
	file.DeleteSheet("Summary")
	file.DeleteSheet(tgtSheet)
	file.NewSheet(tgtSheet)
	var streamWriter *excelize.StreamWriter
	if streamWriter, err = file.NewStreamWriter(tgtSheet); err != nil {
		return
	}
	// write headers
	var style int
	style, err = file.NewStyle(`{
		"font": {
			"size": 10,
			"bold": true
		},
		"fill": {
			"type": "pattern",
			"color": ["#FFFF00"],
			"pattern": 1
		},
		"border": [
			{
				"type": "left",
				"color": "000000",
				"style": 1
			},
			{
				"type": "top",
				"color": "000000",
				"style": 1
			},
			{
				"type": "bottom",
				"color": "000000",
				"style": 1
			},
			{
				"type": "right",
				"color": "000000",
				"style": 1
			}
		],
		"alignment": {
			"horizontal": "center"
		}
	}`)
	if err != nil {
		return
	}
	for i, width := range entity.GetMasterResPlanColumnsWidth() {
		streamWriter.SetColWidth(i+1, i+1, float64(width))
	}
	// first header
	if err = streamWriter.SetRow("A1", []interface{}{excelize.Cell{Value: "Total", StyleID: style}}); err != nil {
		return
	}
	cols := entity.GetMasterResPlanColumns(style)
	cell1, _ := excelize.CoordinatesToCellName(len(cols), 1)
	if err = streamWriter.MergeCell("A1", cell1); err != nil {
		return
	}
	// columns
	if err = streamWriter.SetRow("A2", cols); err != nil {
		return
	}
	cell2, _ := excelize.CoordinatesToCellName(len(cols), 2)
	if err = file.SetCellStyle(tgtSheet, "A2", cell2, style); err != nil {
		return
	}
	// write data
	for rowId, item := range masterResPlanItems {
		cell, _ := excelize.CoordinatesToCellName(1, rowId+3)
		if err = streamWriter.SetRow(cell, item.ConvertToInterfaceArr()); err != nil {
			return
		}
	}
	// save file
	if err = streamWriter.Flush(); err != nil {
		return
	}
	if err = file.SaveAs(dir + "/" + filename); err != nil {
		return
	}
	return filename, nil
}

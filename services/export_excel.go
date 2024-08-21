package services

import (
	"fmt"
	"htmx-learning/entity"
	"log"
	"os"

	_ "image/png"

	"github.com/xuri/excelize/v2"
)

type ExportExcel interface {
	ExportExcelDaily(dailyData []entity.DataTmps) error
}
type exportExcel struct {
	excel *excelize.File
}

func NewExportExcel(excel *excelize.File) ExportExcel {
	return exportExcel{excel}
}

func (e exportExcel) ExportExcelDaily(dailyData []entity.DataTmps) error {
	sheetName := "Sheet1"

	index, _ := e.excel.NewSheet(sheetName)

	// Set table headers
	headers := []string{"Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (PW)", "Q (MVAR)", "PF (%)"}
	for i, header := range headers {
		cell := string(rune('A'+i)) + "7" // A1, B1, C1, etc.
		e.excel.SetCellValue(sheetName, cell, header)
	}

	centerRowStyleId, err := e.excel.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}

	if err := e.excel.SetRowStyle(sheetName, 4, 5, centerRowStyleId); err != nil {
		return err
	}

	Title := "Daily Load Report"
	SubSstationName := "Substation Name"
	bayName := "Bay Name"
	if err := e.excel.SetCellValue(sheetName, "A4", Title); err != nil {
		return err
	}
	if err := e.excel.SetCellValue(sheetName, "A5", SubSstationName); err != nil {
		return err
	}

	if err := e.excel.SetCellValue(sheetName, "A6", fmt.Sprintf("xxx kV %s No.xx", bayName)); err != nil {
		return err
	}

	// Set table data
	// data := [][]interface{}{
	// 	{1, "John Doe", 25, 1},
	// 	{2, "Jane Smith", 30, 2},
	// 	{3, "Sam Brown", 22, 3},
	// }

	// for i, row := range data {
	// 	for j, value := range row {
	// 		cell := string(rune('A'+j)) + fmt.Sprintf("%d", 8+i)
	// 		fmt.Println(cell, value)
	// 		e.excel.SetCellValue(sheetName, cell, value)
	// 	}
	// }
	for row, data := range dailyData {
		cell := "A" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.DataDatetime.Format("02/01/2006"))
		cell = "B" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.DataDatetime.Format("15:04"))
		cell = "C" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, "")
		cell = "D" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, "")
		cell = "E" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, "")
		cell = "F" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.CurrentPhaseA)
		cell = "G" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.CurrentPhaseB)
		cell = "H" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.CurrentPhaseC)
		cell = "I" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.ActivePower)
		cell = "J" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.ReactivePower)
		cell = "K" + fmt.Sprintf("%d", 8+row)
		e.excel.SetCellValue(sheetName, cell, data.PowerFactor)
	}

	if err := e.excel.MergeCell(sheetName, "A4", "K4"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "A5", "K5"); err != nil {
		return err
	}
	if err := e.excel.MergeCell(sheetName, "A6", "K6"); err != nil {
		return err
	}

	// Define the table range
	tableRange := "A7:K26" // Includes headers and data

	// Create a table with the defined range
	disable := true
	if err := e.excel.AddTable(sheetName, &excelize.Table{
		Range:             tableRange,
		Name:              "Table1",
		StyleName:         "TableStyleMedium9",
		ShowFirstColumn:   false,
		ShowLastColumn:    false,
		ShowRowStripes:    &disable,
		ShowColumnStripes: false,
	}); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Set bold borders for the table
	// boldBorderStyle := excelize.Border{
	// 	Type:  "top",
	// 	Color: "000000",
	// 	Style: 2, // Bold border style
	// }
	imagePath := "static/image.png"

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Fatalf("Image file does not exist: %v", err)
		return err
	}
	if err := e.excel.AddPicture(sheetName, "A3", imagePath, &excelize.GraphicOptions{
		AutoFit: false,
		OffsetX: 10,
		OffsetY: 10,
		ScaleX:  0.3,
		ScaleY:  0.3,
	}); err != nil {
		return err
	}

	if err := e.excel.SetColWidth(sheetName, "A", "K", 15); err != nil {
		return err
	}

	stypeId, err := e.excel.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 2},
			{Type: "top", Color: "000000", Style: 2},
			{Type: "bottom", Color: "000000", Style: 2},
			{Type: "right", Color: "000000", Style: 2},
		},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})
	if err != nil {
		return err
	}
	// Apply bold borders to all sides
	err = e.excel.SetCellStyle(sheetName, "A6", "K54", stypeId)
	if err != nil {
		return err
	}

	// Set active sheet and save the file
	e.excel.SetActiveSheet(index)

	if err = e.excel.SaveAs("Book1.xlsx"); err != nil {
		return err
	}

	return nil
}

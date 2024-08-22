package services

import (
	"fmt"
	"htmx-learning/entity"

	"github.com/jung-kurt/gofpdf"
)

func ExportPdfDaily(dailyData []entity.DataTmps, fileName string) error {

	filePath := fileName

	substationName := "Substation Name"
	bayName := "XXX kV Bay Name No. XX"
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	imagePath := "static/image.png"
	pdf.Image(imagePath, 10, 15, 20, 0, false, "", 0, "")
	// Set font
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(187, 10, "Daily Load Report", "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, substationName, "", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.CellFormat(187, 10, bayName, "LTRB", 0, "C", false, 0, "")
	pdf.SetY(40)

	// Header
	headers := []string{"Date", "Time", "Vab (kV)", "Vbc (kV)", "Vca (kV)", "Ia (A)", "Ib (A)", "Ic (A)", "P (MW)", "Q (MVAR)", "PF (%)"}
	for _, header := range headers {
		pdf.CellFormat(17, 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Table rows
	pdf.SetFont("Arial", "", 8)

	for _, data := range dailyData {
		pdf.CellFormat(17, 10, data.DataDatetime.Format("02/01/2006"), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, data.DataDatetime.Format("15:04"), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, "", "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseA), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseB), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.CurrentPhaseC), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ActivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.ReactivePower), "1", 0, "C", false, 0, "")
		pdf.CellFormat(17, 10, fmt.Sprintf("%.2f", data.PowerFactor), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	// Save the PDF to a file
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return err
	}
	return nil
}

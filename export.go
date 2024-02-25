package main

import (
	"log"
	"time"

	"github.com/go-pdf/fpdf"
)

func createPDF(fn string, timeframe string) {
	var data []string
	if timeframe == "month" {
		data = monthReport(fn, settings)
	} else {
		data = weekReport(fn, settings)
	}
	day := time.Now().Format(time.DateOnly)
	outfile := day + ".pdf"
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	leftMargin, _, rightMargin, _ := pdf.GetMargins()
	pageWidth, _ := pdf.GetPageSize()
	pageWidth -= leftMargin + rightMargin
	for _, element := range data {
		// X, Y, TEXT
		pdf.WriteAligned(pageWidth, 35, element, "L")
		pdf.Ln(5)
	}
	err := pdf.OutputFileAndClose(outfile)
	if err != nil {
		log.Fatal(err)
	}
}

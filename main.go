package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	// Generate barcode
	content := "One Piece Is The Best Anime"
	barcode, err := generateBarcode(content)
	if err != nil {
		log.Fatalf("Error generating barcode: %s", err)
	}

	filenamePrefix := "onepiece"
	filename, err := generateBarcodeFile(barcode, filenamePrefix)
	if err != nil {
		log.Fatalf("Error creating barcode file: %s", err)
	}

	// Create pdf
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	// MAWB No
	drawTextWithNewline(0, 0, 30, "MAWB No: 131-04455301", pdf)
	drawImage(9, 20, 0, 15, filename, pdf)

	// HAWB No
	drawTextWithNewline(0, 10, 30, "HAWB No: XASUZ2407476", pdf)
	drawImage(9, 55, 0, 15, filename, pdf)

	// Shipper
	drawTextWithNewline(0, 20, 30, "Shipper: JIANGSU DINGS' INTELLIGENT CONTROL TECHNOLOGY CO., LTD", pdf)
	drawImage(9, 90, 0, 15, filename, pdf)

	// No. of package(s)
	drawTextWithNewline(0, 30, 30, "No. of package(s): 1", pdf)
	drawImage(9, 125, 0, 15, filename, pdf)

	// Gross weight (kg)
	drawTextWithNewline(0, 40, 30, "Gross weight (kg): 94.00", pdf)
	drawImage(9, 160, 0, 15, filename, pdf)

	// Gross weight (kg)
	drawTextWithNewline(0, 50, 30, "Chargeable weight (kg): 94.00", pdf)
	drawImage(9, 195, 0, 15, filename, pdf)

	err = pdf.OutputFileAndClose("/mnt/c/Users/AWOT/Desktop/test.pdf")
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}

}

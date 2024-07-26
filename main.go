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
	drawTextWithNewline(0, 0, 30, "MAWB No: 000-00000000", pdf)
	drawImage(9, 20, 0, 15, filename, pdf)

	// HAWB No
	drawTextWithNewline(0, 10, 30, "HAWB No: 12345", pdf)
	drawImage(9, 55, 0, 15, filename, pdf)

	// Shipper
	drawTextWithNewline(0, 20, 30, "Shipper: Some Shipper", pdf)
	drawImage(9, 90, 0, 15, filename, pdf)

	// No. of package(s)
	drawTextWithNewline(0, 30, 30, "No. of package(s): 1", pdf)
	drawImage(9, 125, 0, 15, filename, pdf)

	// Gross weight (kg)
	drawTextWithNewline(0, 40, 30, "Gross weight (kg): 1.00", pdf)
	drawImage(9, 160, 0, 15, filename, pdf)

	// Gross weight (kg)
	drawTextWithNewline(0, 50, 30, "Chargeable weight (kg): 1.00", pdf)
	drawImage(9, 195, 0, 15, filename, pdf)

	err = pdf.OutputFileAndClose("/mnt/c/Users/AWOT/Desktop/test.pdf")
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}

}

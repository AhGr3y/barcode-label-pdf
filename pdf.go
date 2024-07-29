package main

import (
	"log"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

// drawTextWithNewline - Draws <content> on the current pdf page
// on (<x>, <y>) coordinate based on the current pointer location.
// Then creates a newline with <newlineHeight> below <content>.
// This sets the pointer location directly below the newline with
// its x-coordinate set to 0. If the there is a page margin, the
// pointer will be placed right after the margin.
func drawTextWithNewline(x, y, newlineHeight float64, content string, pdf *gofpdf.Fpdf) {
	pdf.Cell(x, y, content)
	pdf.Ln(newlineHeight)
}

// drawImage - Draws the image in the given file using the path
// <filename> at the given (<x>, <y>) coordinates. Sets the width
// and height of the image at <w> and <h>.
func drawImage(x, y, w, h float64, filename string, pdf *gofpdf.Fpdf) {
	pdf.ImageOptions(filename, x, y, w, h, false, gofpdf.ImageOptions{ImageType: "JPEG", ReadDpi: true}, 0, "")
}

func generatePDF(inputs []string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	prefixes := []string{
		"MAWB No: ",
		"HAWB No: ",
		"Shipper: ",
		"No. of package(s): ",
		"Gross weight (kg): ",
		"Chargeable weight (kg): ",
	}

	for i, input := range inputs {

		barcode, err := generateBarcode(input)
		if err != nil {
			log.Fatalf("Error generating barcode: %s", err)
		}

		fileBasename := strconv.Itoa(i)
		filePathname, err := generateBarcodeFile(barcode, fileBasename)
		if err != nil {
			log.Fatalf("Error creating barcode file: %s", err)
		}

		text := prefixes[i] + input
		imageY := float64(20 + (35 * i))
		drawTextWithNewline(0, float64(10*i), 30, text, pdf)
		drawImage(9, imageY, 0, 15, filePathname, pdf)

	}

	err := pdf.OutputFileAndClose("/mnt/c/Users/AWOT/Desktop/test.pdf")
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}

	return nil
}

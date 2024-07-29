package main

import (
	"log"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	contents := []string{
		"131-12345678",
		"XASUZ2407000",
		"JIANGSU' TECHNOLOGY CO., LTD",
		"1",
		"70.00",
		"100.00",
	}

	prefixes := []string{
		"MAWB No: ",
		"HAWB No: ",
		"Shipper: ",
		"No of package(s): ",
		"Gross weight (kg): ",
		"Chargeable weight (kg): ",
	}

	// Create pdf
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	for i, content := range contents {

		barcode, err := generateBarcode(content)
		if err != nil {
			log.Fatalf("Error generating barcode: %s", err)
		}

		fileBasename := strconv.Itoa(i)
		filePathname, err := generateBarcodeFile(barcode, fileBasename)
		if err != nil {
			log.Fatalf("Error creating barcode file: %s", err)
		}

		text := prefixes[i] + content
		imageY := float64(20 + (35 * i))
		drawTextWithNewline(0, float64(10*i), 30, text, pdf)
		drawImage(9, imageY, 0, 15, filePathname, pdf)

	}

	err := pdf.OutputFileAndClose("/mnt/c/Users/AWOT/Desktop/test.pdf")
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}

}

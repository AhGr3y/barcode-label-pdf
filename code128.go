package main

import (
	"image/jpeg"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

// generateBarcode - Uses <content> to generate a code128 barcode.
func generateBarcode(content string) (barcode.Barcode, error) {

	// Set width and height
	width := len(content) * 17
	if width > 900 {
		width = 900
	}
	height := 50

	// Create the barcode
	bc, err := code128.Encode(content)
	if err != nil {
		return nil, err
	}

	// Scale the barcode to width x height pixels
	bcode, err := barcode.Scale(bc, width, height)
	if err != nil {
		return nil, err
	}

	return bcode, nil
}

// generateBarcodeFile - Uses <barcode> to generate a <filename>.png file
// in the ./output directory. Returns the relative path to the file from the
// root directory if file successfully created.
func generateBarcodeFile(barcode barcode.Barcode, filename string) (string, error) {

	filename = "./output/" + filename + ".jpeg"

	// create the output file
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// encode the barcode as png
	err = jpeg.Encode(file, barcode, nil)
	if err != nil {
		return "", err
	}

	return filename, nil
}

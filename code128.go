package main

import (
	"errors"
	"image/jpeg"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

// generateBarcode - Uses <content> to generate a code128 barcode.
func generateBarcode(content string) (barcode.Barcode, error) {

	// Handling edge cases
	if content == "" || len(content) > 80 {
		return nil, errors.New("content length must be between 1 and 79 inclusive")
	}

	// Create the barcode
	bc, err := code128.Encode(content)
	if err != nil {
		return nil, err
	}

	// Set width to smallest possible length for
	// the generated barcode and height to 50px.
	bcBounds := bc.Bounds()
	width := bcBounds.Max.X - bcBounds.Min.X
	height := 50

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
func generateBarcodeFile(barcode barcode.Barcode, fileBasename string) (string, error) {

	filePathname := "./output/" + fileBasename + ".jpeg"

	// create the output file
	file, err := os.Create(filePathname)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// encode the barcode as png
	err = jpeg.Encode(file, barcode, nil)
	if err != nil {
		return "", err
	}

	return filePathname, nil
}

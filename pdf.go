package main

import "github.com/jung-kurt/gofpdf"

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

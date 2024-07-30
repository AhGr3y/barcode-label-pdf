package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func runGUI() {
	a := app.New()
	w := a.NewWindow("Barcode PDF Generator")
	w.Resize(fyne.NewSize(800, 550))

	showHomepage(w)
	//showCode128Page(w)

	w.ShowAndRun()
}

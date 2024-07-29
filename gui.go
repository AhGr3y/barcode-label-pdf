package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func runGUI() {
	a := app.New()
	w := a.NewWindow("Barcode PDF Generator")
	w.Resize(fyne.NewSize(800, 550))

	// Logger for displaying information to users
	logger := widget.NewLabel("> Welcome to Barcode PDF Generator!")

	// Entry widgets for each form field
	mawbEntry := widget.NewEntry()
	hawbEntry := widget.NewEntry()
	shipperEntry := widget.NewEntry()
	pkgsEntry := widget.NewEntry()
	gwEntry := widget.NewEntry()
	cwEntry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "", Widget: &fyne.Container{}},
			{Text: "MAWB No:", Widget: mawbEntry},
			{Text: "", Widget: &fyne.Container{}},
			{Text: "HAWB No:", Widget: hawbEntry},
			{Text: "", Widget: &fyne.Container{}},
			{Text: "Shipper:", Widget: shipperEntry},
			{Text: "", Widget: &fyne.Container{}},
			{Text: "No of package(s):", Widget: pkgsEntry},
			{Text: "", Widget: &fyne.Container{}},
			{Text: "Gross weight (kg):", Widget: gwEntry},
			{Text: "", Widget: &fyne.Container{}},
			{Text: "Chargeable weight (kg):", Widget: cwEntry},
		},
		SubmitText: "Generate PDF",
		OnSubmit: func() {

			inputs := []string{
				mawbEntry.Text,
				hawbEntry.Text,
				shipperEntry.Text,
				pkgsEntry.Text,
				gwEntry.Text,
				cwEntry.Text,
			}
			err := generatePDF(inputs)
			now := time.Now().Format(time.DateTime)
			if err != nil {
				logger.SetText("> " + now + ": " + err.Error())
			}
			if err == nil {
				logger.SetText("> " + now + ": PDF Generated!")
			}
		},
		CancelText: "Exit Application",
		OnCancel: func() {
			w.Close()
		},
	}

	w.SetContent(
		container.NewBorder(
			nil,
			container.NewBorder(
				widget.NewSeparator(),
				nil,
				nil,
				nil,
				logger,
			),
			nil,
			nil,
			form,
		),
	)

	w.ShowAndRun()
}

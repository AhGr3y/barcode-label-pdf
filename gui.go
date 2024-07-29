package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func runGUI() {
	a := app.New()
	w := a.NewWindow("MyApp")
	w.Resize(fyne.NewSize(300, 400))

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
			if err != nil {
				log.Printf("Error generating pdf: %s", err)
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
			nil,
			nil,
			nil,
			form,
		),
	)

	w.ShowAndRun()
}

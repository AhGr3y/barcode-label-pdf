package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func showHomepage(w fyne.Window) {

	titleLabel := widget.NewLabelWithStyle(
		"Welcome to Barcode PDF Generator!",
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Bold: true,
		},
	)

	helperLabel := widget.NewLabelWithStyle(
		"View your templates here:",
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Italic: true,
		},
	)

	exitButton := widget.NewButton("Exit", func() { w.Close() })

	templatesButton := widget.NewButton("Templates", func() { showTemplatesPage(w) })

	content := container.NewCenter(
		container.NewVBox(
			titleLabel,
			widget.NewLabel(""),
			helperLabel,
			templatesButton,
			widget.NewLabel("\n"),
			exitButton,
		),
	)

	w.SetContent(content)
}

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
		"Select one of the formats below to get started:",
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Italic: true,
		},
	)

	exitButton := widget.NewButton("Exit", func() { w.Close() })

	code128Button := widget.NewButton("Code128", func() { showCode128Page(w) })

	content := container.NewCenter(
		container.NewVBox(
			titleLabel,
			widget.NewLabel(""),
			helperLabel,
			code128Button,
			widget.NewLabel("\n"),
			exitButton,
		),
	)

	w.SetContent(content)
}

package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

const (
	CODE128 = "code128"
	QRCODE  = "qrcode"
)

func showHomepage(w fyne.Window) {

	titleLabel := widget.NewLabelWithStyle(
		"Welcome to Barcode PDF Generator!",
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Bold: true,
		},
	)

	selectTemplatesLabel := widget.NewLabelWithStyle(
		"View your templates here:",
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Italic: true,
		},
	)

	createTemplateLabel := widget.NewLabelWithStyle(
		"Or, create a template using one of these formats:",
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Italic: true,
		},
	)

	exitButton := widget.NewButton("Exit", func() { w.Close() })
	templatesButton := widget.NewButton("Templates", func() { showTemplatesPage(w) })
	createCode128Button := widget.NewButton("Code128", func() { showCreateTemplatePage(CODE128, w) })

	content := container.NewCenter(
		container.NewVBox(
			titleLabel,
			widget.NewLabel(""),
			selectTemplatesLabel,
			templatesButton,
			widget.NewLabel(""),
			createTemplateLabel,
			createCode128Button,
			widget.NewLabel("\n"),
			exitButton,
		),
	)

	w.SetContent(content)
}

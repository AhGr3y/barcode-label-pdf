package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func showCreateTemplatePage(codeType string, w fyne.Window) {

	logger := widget.NewLabel("> Fill in the form to create the corresponding field.")

	if codeType == "" {
		logger.SetText("> Code type not specified, will default to code128 format.")
		codeType = CODE128
	}

	templates := container.NewCenter(
		container.NewVBox(
			widget.NewLabel("Entry for creating fields goes here..."),
		),
	)

	content := container.NewBorder(
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
		templates,
	)

	w.SetContent(content)
}

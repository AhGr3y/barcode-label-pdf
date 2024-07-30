package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func showCode128Page(fields []string, w fyne.Window) {
	// Logger for displaying information to users
	logger := widget.NewLabel("> Fill in the fields to generate a PDF.")

	entries := []*widget.Entry{}

	for range fields {
		entry := widget.NewEntry()
		entry.SetPlaceHolder("Enter text here...")
		entries = append(entries, entry)
	}

	items := []*widget.FormItem{}

	for i, entry := range entries {
		items = append(items, &widget.FormItem{
			Text:   "",
			Widget: &fyne.Container{},
		})
		items = append(items, &widget.FormItem{
			Text:   fields[i] + ":",
			Widget: entry,
		})
	}

	form := &widget.Form{
		Items:      items,
		SubmitText: "Generate PDF",
		OnSubmit: func() {
			inputs := []string{}
			prefixes := []string{}

			for i, entry := range entries {
				inputs = append(inputs, entry.Text)
				prefixes = append(prefixes, fields[i]+": ")
			}

			filepath, err := generatePDF(inputs, prefixes)
			now := time.Now().Format(time.DateTime)
			if err != nil {
				logger.SetText("> " + now + ": " + err.Error())
			}
			if err == nil {
				logger.SetText("> " + now + ": PDF generated at " + filepath)
			}
		},
		CancelText: "Back to Homepage",
		OnCancel: func() {
			showHomepage(w)
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
}

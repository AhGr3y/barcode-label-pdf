package main

import (
	"errors"
	"os"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func showCreateTemplatePage(codeType string, w fyne.Window) {

	timeNow := time.Now().Format(time.DateTime)
	logger := widget.NewLabel("> " + timeNow + ": Fill in the form to create the corresponding field.")

	if codeType == "" {
		logger.SetText("> " + timeNow + ": Code type not specified, will default to code128 format.")
		codeType = CODE128
	}

	cardContainer := container.NewVBox()
	addedFieldCard := widget.NewCard("", "", cardContainer)

	enteredFieldsStr := []string{}
	enteredFieldsObj := []fyne.CanvasObject{}
	fieldEntry := widget.NewEntry()
	fieldEntry.SetPlaceHolder("Enter field name, e.g. First Name")
	fieldEntry.Validator = func(s string) error {
		if len(s) == 0 {
			return errors.New("> " + timeNow + ": Field cannot be empty.")
		}

		return nil
	}

	fieldLabel := widget.NewLabelWithStyle(
		"List of fields added:",
		fyne.TextAlignCenter,
		fyne.TextStyle{
			Italic: true,
		},
	)

	fieldButton := widget.NewButton("Add Field", func() {
		err := fieldEntry.Validate()
		if err != nil {
			logger.SetText(err.Error())
			return
		}
		label := widget.NewLabel(fieldEntry.Text)
		enteredFieldsObj = append(enteredFieldsObj, label)
		enteredFieldsStr = append(enteredFieldsStr, fieldEntry.Text)
		cardContainer.Add(label)
		fieldEntry.SetText("")
		logger.SetText("> " + timeNow + ": Added field.")
	})

	undoButton := widget.NewButton("Undo", func() {
		if len(enteredFieldsObj) == 0 {
			logger.SetText("> " + timeNow + ": Nothing to undo.")
			return
		}
		cardContainer.Remove(enteredFieldsObj[len(enteredFieldsObj)-1])
		logger.SetText("> " + timeNow + ": Removed field.")
		enteredFieldsObj = enteredFieldsObj[:len(enteredFieldsObj)-1]
		enteredFieldsStr = enteredFieldsStr[:len(enteredFieldsStr)-1]
	})

	entryContainer := container.NewHBox(
		fieldEntry,
		fieldButton,
		undoButton,
	)

	saveEntry := widget.NewEntry()
	saveEntry.SetPlaceHolder("Enter template name...")
	saveEntry.Validator = func(s string) error {
		if saveEntry.Text == "" {
			return errors.New("> " + timeNow + ": Cannot have empty template name.")
		}

		if len(enteredFieldsStr) == 0 {
			return errors.New("> " + timeNow + ": Cannot create template without adding field(s).")
		}

		specialCharacters := []string{
			"!", "@", "#", "$", "%", "`", "~", "^", "&", "*", "(", ")", "_", "-", "+", "=",
		}
		for _, c := range specialCharacters {
			if strings.Contains(s, c) {
				return errors.New("> " + timeNow + ": Filename cannot contain special characters.")
			}
		}

		return nil
	}

	saveButton := widget.NewButton("Save Template", func() {
		err := saveEntry.Validate()
		if err != nil {
			logger.SetText(err.Error())
			return
		}

		dataString := ""
		for _, field := range enteredFieldsStr {
			dataString += field + "\n"
		}
		data := []byte(dataString)
		filename := codeType + "_" + saveEntry.Text

		err = os.WriteFile("./templates/"+filename, data, 0644)
		if err != nil {
			logger.SetText("> " + timeNow + ": Something went wrong, please try again.")
		}

		logger.SetText("> " + timeNow + ": Template created.")
	})

	saveContainer := container.NewCenter(
		container.NewHBox(
			saveEntry,
			saveButton,
		),
	)

	homeButton := widget.NewButton("Back to Homepage", func() { showHomepage(w) })

	templates := container.NewCenter(
		container.NewVBox(
			fieldLabel,
			addedFieldCard,
			entryContainer,
			widget.NewLabel("\n"),
			saveContainer,
			widget.NewLabel("\n"),
			homeButton,
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

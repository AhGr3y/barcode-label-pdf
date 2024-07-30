package main

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"os"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func getTemplateNameFromFilename(filename string) (string, error) {

	if filename == "" {
		return "", errors.New("unable to get template name from empty string")
	}

	split := strings.Split(filename, "_")
	return split[1], nil
}

func createButtonFromTemplate(templateFile fs.DirEntry, w fyne.Window) *widget.Button {
	fields := []string{}

	file, err := os.Open("./templates/" + templateFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	templateName, err := getTemplateNameFromFilename(templateFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields = append(fields, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	button := widget.NewButton(templateName, func() {
		showCode128Page(fields, w)
	})

	return button
}

func showTemplatesPage(w fyne.Window) {

	logger := widget.NewLabel("> Select a template.")

	templateFiles, err := os.ReadDir("./templates")
	if err != nil {
		log.Fatal(err)
	}

	if len(templateFiles) == 0 {
		logger.SetText("> There are no templates, please create a template to get started.")
	}

	buttons := []fyne.CanvasObject{}

	// Create a button for each templateFile
	for _, templateFile := range templateFiles {
		buttons = append(buttons, createButtonFromTemplate(templateFile, w))
	}

	templates := container.NewVBox(
		buttons...,
	)

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
			templates,
		),
	)
}

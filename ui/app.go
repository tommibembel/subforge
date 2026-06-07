package ui

import (
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/tommibembel/subforge/internal/parser"
	"github.com/tommibembel/subforge/internal/writer"
)

type App struct {
	fyneApp fyne.App
	window  fyne.Window

	inputPath  string
	outputPath string

	inputLabel  *widget.Label
	outputLabel *widget.Label
	formatSelect *widget.Select
	statusLabel *widget.Label
}

func Run() {
	a := app.NewWithID("com.bendel-translations.subforge")
	w := a.NewWindow("SubForge")
	w.Resize(fyne.NewSize(700, 400))
	w.SetContent(buildUI(a, w))
	w.ShowAndRun()
}

func buildUI(a fyne.App, w fyne.Window) fyne.CanvasObject {
	inputLabel := widget.NewLabel("No file selected")
	inputLabel.Wrapping = fyne.TextTruncate

	outputLabel := widget.NewLabel("No file selected")
	outputLabel.Wrapping = fyne.TextTruncate

	statusLabel := widget.NewLabel("")

	formats := []string{"docx", "xlsx", "txt", "flow"}
	formatSelect := widget.NewSelect(formats, nil)
	formatSelect.SetSelected("docx")

	var inputPath, outputPath string

	selectInput := widget.NewButton("Select input file", func() {
		dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
			if err != nil || f == nil {
				return
			}
			inputPath = f.URI().Path()
			inputLabel.SetText(inputPath)

			ext := formatSelect.Selected
			if ext == "flow" {
				ext = "txt"
			}
			outputPath = strings.TrimSuffix(inputPath, ".srt") + "." + ext
			outputLabel.SetText(outputPath)
		}, w)
	})

	selectOutput := widget.NewButton("Select output file", func() {
		dialog.ShowFileSave(func(f fyne.URIWriteCloser, err error) {
			if err != nil || f == nil {
				return
			}
			outputPath = f.URI().Path()
			outputLabel.SetText(outputPath)
		}, w)
	})

	formatSelect.OnChanged = func(selected string) {
		if inputPath == "" || outputPath == "" {
			return
		}
		ext := selected
		if ext == "flow" {
			ext = "txt"
		}
		base := strings.TrimSuffix(inputPath, filepath.Ext(inputPath))
		outputPath = base + "." + ext
		outputLabel.SetText(outputPath)
	}

	convertBtn := widget.NewButtonWithIcon("Convert", theme.ConfirmIcon(), func() {
		if inputPath == "" || outputPath == "" {
			statusLabel.SetText("Please select input and output file")
			return
		}
		statusLabel.SetText("Converting...")

		subtitles, err := parser.ParseSRT(inputPath)
		if err != nil {
			statusLabel.SetText("Error: " + err.Error())
			return
		}

		switch formatSelect.Selected {
		case "docx":
			err = writer.WriteDOCX(outputPath, subtitles)
		case "xlsx":
			err = writer.WriteXLSX(outputPath, subtitles)
		case "txt":
			err = writer.WriteTXT(outputPath, subtitles)
		case "flow":
			err = writer.WriteTXTFlowing(outputPath, subtitles)
		}

		if err != nil {
			statusLabel.SetText("Error: " + err.Error())
			return
		}

		statusLabel.SetText("Conversion successful: " + outputPath)
		inputPath = ""
		outputPath = ""
		inputLabel.SetText("No file selected")
		outputLabel.SetText("No file selected")
	})

	content := container.NewVBox(
		widget.NewLabel("Input file:"),
		inputLabel,
		selectInput,
		widget.NewSeparator(),
		widget.NewLabel("Output format:"),
		formatSelect,
		widget.NewSeparator(),
		widget.NewLabel("Output file:"),
		outputLabel,
		selectOutput,
		widget.NewSeparator(),
		convertBtn,
		statusLabel,
	)

	return container.NewPadded(content)
}

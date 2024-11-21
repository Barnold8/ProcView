package main

import (
	"image/color"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type ProcWindow struct {
	x               float32
	y               float32
	width           float32
	height          float32
	app             fyne.App
	window          fyne.Window
	window_contents *fyne.Container
}

func CreateBox() *fyne.Container {
	// Box background
	rect := canvas.NewRectangle(color.NRGBA{R: 200, G: 200, B: 255, A: 255}) // Light blue background
	rect.SetMinSize(fyne.NewSize(0, 30))                                     // Set minimum height for boxes

	// Label for text
	label := widget.NewLabel("")
	label.Alignment = fyne.TextAlignCenter

	// Combine background and label into a single container
	return container.NewStack(rect, label)
}

func AppendData(previous_data binding.ExternalStringList, processes map[string]Process) {

	for {

		processData := ProcessMapToString(UpdateProcesses(processes, time.Now(), string(grabProcesses())))
		// fmt.Println(processData)

		err := previous_data.Set(append([]string{"Name, Start, Time"}, strings.Split(processData, "\n")...))

		if err != nil {

			panic(err)
		}
	}
}

package main

import (
	"image/color"
	"strings"
	"sync"
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

var (
	mu           sync.Mutex    // Mutex for thread-safe access to the global variable
	controlValue EControlValue // Global integer that will control the behavior of processData
)

func CreateBox() *fyne.Container {
	// Box background
	rect := canvas.NewRectangle(color.NRGBA{R: 40, G: 41, B: 46, A: 255}) // Light blue background
	rect.SetMinSize(fyne.NewSize(0, 30))                                  // Set minimum height for boxes

	// Label for text
	label := widget.NewLabel("")
	label.Alignment = fyne.TextAlignCenter

	// Combine background and label into a single container
	return container.NewStack(rect, label)
}

func AppendData(previous_data binding.ExternalStringList, processes map[string]Process, list *widget.List) {

	for {

		var processData string

		mu.Lock()

		switch controlValue {

		case 0:
			{
				processData = ProcessMapToStringSortedByName(UpdateProcesses(processes, time.Now(), string(grabProcesses())), false)
				break
			}

		case 1:
			{
				processData = ProcessMapToStringSortedByName(UpdateProcesses(processes, time.Now(), string(grabProcesses())), true)
				break
			}

		case 2:
			{
				processData = ProcessMapToStringSortedByTimeAlive(UpdateProcesses(processes, time.Now(), string(grabProcesses())), false)
				break
			}

		case 3:
			{
				processData = ProcessMapToStringSortedByTimeAlive(UpdateProcesses(processes, time.Now(), string(grabProcesses())), true)
				break
			}

		case 4:
			{
				processData = ProcessMapToStringSortedByTimeStarted(UpdateProcesses(processes, time.Now(), string(grabProcesses())), false)
				break
			}

		case 5:
			{
				processData = ProcessMapToStringSortedByTimeStarted(UpdateProcesses(processes, time.Now(), string(grabProcesses())), true)
				break
			}

		default:

			processData = ProcessMapToStringSortedByName(UpdateProcesses(processes, time.Now(), string(grabProcesses())), false)

		}

		err := previous_data.Set(strings.Split(processData, "\n"))

		if err != nil {
			panic(err)
		}
		if list != nil {
			list.Refresh()
		}

		mu.Unlock()
	}
}

func NameSignal() {

	mu.Lock()

	if controlValue != ByName && controlValue != ByNameInverse {
		controlValue = ByName
	} else {
		if controlValue == ByName {
			controlValue = ByNameInverse
		} else {
			controlValue = ByName
		}
	}

	mu.Unlock()

}

func TimeAliveSignal() {

	mu.Lock()

	if controlValue != ByAlive && controlValue != ByAliveInverse {
		controlValue = ByAliveInverse
	} else {
		if controlValue == ByAlive {
			controlValue = ByAliveInverse
		} else {
			controlValue = ByAlive
		}
	}

	mu.Unlock()

}

func TimeCreatedSignal() {

	mu.Lock()

	if controlValue != ByCreated && controlValue != ByCreatedInverse { // Not set in this range
		controlValue = ByCreated
	} else {
		if controlValue == ByCreated {
			controlValue = ByCreatedInverse
		} else {
			controlValue = ByCreated
		}
	}

	mu.Unlock()

}

func File(value string, combo *widget.Select) {

	if combo != nil {
		combo.Selected = "File"
	}

	switch value {
	case "Export CSV":
		SaveToFile("Export.csv", ProcessMapToCSV(UpdateProcesses(ParseProcesses(string(grabProcesses())), time.Now(), string(grabProcesses()))))

	default:

	}

}

func Settings(value string, combo *widget.Select) {

	if combo != nil {
		combo.Selected = "Settings"
	}

	switch value {
	case "Add To StartUp":

		break

	default:

	}

}

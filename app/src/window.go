package main

import (
	"fmt"
	"image/color"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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
		err := addProgramToStartup("ProcView.exe")
		if err != nil {
			fmt.Printf("Error detected when adding program to startup %sn", err)
		}
		break

	case "Remove From StartUp":
		err := removeFromStartUp()
		if err != nil {
			fmt.Printf("Error detected when adding program to startup %sn", err)
		}
		break

	default:

	}

}
func makeWindow() {

	// variable definitions/declarations
	var fileCombo *widget.Select
	var settingsCombo *widget.Select

	myFunc := AppendData

	data := binding.BindStringList(
		&[]string{ProcessMapToStringSortedByName(ParseProcesses(string(grabProcesses())), true)},
	)
	// variable definitions/declarations

	// GUI MESS, NOT SURE HOW TO MAKE CLEAN

	fileCombo = widget.NewSelect([]string{"Export CSV"}, func(value string) {
		File(value, fileCombo)
	})

	settingsCombo = widget.NewSelect([]string{"Add To StartUp", "Remove From StartUp"}, func(value string) {
		Settings(value, settingsCombo)
	})

	fileCombo.Selected = "File"
	settingsCombo.Selected = "Settings"

	toolbar := container.NewGridWithColumns(2,
		container.New(layout.NewStackLayout(), fileCombo),
		container.New(layout.NewStackLayout(), settingsCombo),
	)

	categories := container.NewGridWithColumns(3, // Four equal columns
		widget.NewButton("Name", NameSignal),
		widget.NewButton("Time created", TimeCreatedSignal),
		widget.NewButton("Runtime", TimeAliveSignal),
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {

			return container.NewGridWithColumns(3,
				CreateBox(), CreateBox(), CreateBox(),
			)
		},
		func(i binding.DataItem, o fyne.CanvasObject) {

			grid := o.(*fyne.Container)
			stringValue, err := i.(binding.String).Get()
			if err != nil {
				return
			}

			labels := strings.Split(stringValue, ",")
			for idx, child := range grid.Objects {
				if idx < len(labels) {
					box := child.(*fyne.Container)
					label := box.Objects[1].(*widget.Label)
					label.SetText(labels[idx])
				}
			}
		})

	// // Create a rectangle for the background
	background := canvas.NewRectangle(color.NRGBA{R: 40, G: 41, B: 46, A: 255}) // Light blue background

	// // Use a vertical box layout for the content
	content := container.NewBorder(categories, nil, nil, nil, list)
	content2 := container.NewBorder(toolbar, nil, nil, nil, content)

	// // Combine background and content using container.NewMax
	mainContent := container.NewStack(
		background,
		content2,
	)

	// GUI MESS, NOT SURE HOW TO MAKE CLEAN

	windowBuilder := ConcreteWindowBuilder{}

	pWindow := windowBuilder.InitialiseWindow().SetWindowContainer(mainContent).SetWindowSize(900, 500).Build()

	pWindow.app.Settings().SetTheme(theme.DarkTheme())

	pWindow.window.SetTitle("ProcView")

	go myFunc(data, ParseProcesses(string(grabProcesses())), list)

	pWindow.window.ShowAndRun()
}

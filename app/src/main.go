package main

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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

	go myFunc(data, ParseProcesses(string(grabProcesses())), list)

	pWindow.window.ShowAndRun()
}

func main() {
	makeWindow()
}

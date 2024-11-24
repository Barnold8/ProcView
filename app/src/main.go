package main

import (
	"image/color"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myFunc := AppendData

	data := binding.BindStringList(
		&[]string{ProcessMapToStringSortedByName(ParseProcesses(string(grabProcesses())), true)},
	)

	// GUI MESS, NOT SURE HOW TO MAKE CLEAN

	combo1 := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Combo 1 set to", value)
	})

	combo2 := widget.NewSelect([]string{"Choice A", "Choice B"}, func(value string) {
		log.Println("Combo 2 set to", value)
	})

	combo3 := widget.NewSelect([]string{"Select X", "Select Y"}, func(value string) {
		log.Println("Combo 3 set to", value)
	})

	// Create a horizontal container
	toolbar := container.NewGridWithColumns(3,
		container.New(layout.NewStackLayout(), combo1),
		container.New(layout.NewStackLayout(), combo2),
		container.New(layout.NewStackLayout(), combo3),
	)

	categories := container.NewGridWithColumns(3, // Four equal columns
		widget.NewButton("Name", NameSignal),
		widget.NewButton("Time created", TimeCreatedSignal),
		widget.NewButton("Runtime", TimeAliveSignal),
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			// Create a row with dynamically sized boxes
			return container.NewGridWithColumns(3, // Adjust column count as necessary
				CreateBox(), CreateBox(), CreateBox(),
			)
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			// Bind data to each box
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

	// // Create layout content
	title := widget.NewLabel("Process Viewer")
	title.Alignment = fyne.TextAlignCenter

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

package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myFunc := AppendData

	data := binding.BindStringList(
		&[]string{ProcessMapToStringSortedByName(ParseProcesses(string(grabProcesses())), true)},
	)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			fmt.Println("Add clicked")
		}),

		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			fmt.Println("Remove clicked")
		}),

		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			fmt.Println("Refresh clicked")
		}),
	)

	categories := container.NewGridWithColumns(3, // Four equal columns
		widget.NewButton("Name", func() {}),
		widget.NewButton("Time created", func() {}),
		widget.NewButton("Runtime", func() {}),
	)

	// GUI MESS, NOT SURE HOW TO MAKE CLEAN
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
	background := canvas.NewRectangle(color.NRGBA{R: 200, G: 200, B: 255, A: 255}) // Light blue background

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

	go myFunc(data, ParseProcesses(string(grabProcesses())), list)

	pWindow.window.ShowAndRun()

}

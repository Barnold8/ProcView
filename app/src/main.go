package main

import (
	"fmt"
	"image/color"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func appendData(previous_data binding.ExternalStringList, processes map[string]Process) {

	for {

		processData := ProcessMapToString(UpdateProcesses(processes, time.Now(), string(grabProcesses())))
		fmt.Println(processData)

		err := previous_data.Set(append([]string{"Name, Start, Time"}, strings.Split(processData, "\n")...))

		if err != nil {

			panic(err)
		}
	}
}

func main() {

	myFunc := appendData

	data := binding.BindStringList(
		&[]string{ProcessMapToString(ParseProcesses(string(grabProcesses())))},
	)

	go myFunc(data, ParseProcesses(string(grabProcesses())))

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
	content := container.NewBorder(title, nil, nil, nil, list)

	// // Combine background and content using container.NewMax
	mainContent := container.NewStack(
		background,
		content,
	)

	// GUI MESS, NOT SURE HOW TO MAKE CLEAN

	windowBuilder := ConcreteWindowBuilder{}

	pWindow := windowBuilder.InitialiseWindow().SetWindowContainer(mainContent).SetWindowSize(900, 500).Build()

	pWindow.window.ShowAndRun()

}

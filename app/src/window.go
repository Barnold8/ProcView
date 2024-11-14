package main

import (
	"fyne.io/fyne/v2"
)

type procWindow struct {
	x         float32
	y         float32
	width     float32
	height    float32
	title     string
	dark_mode bool
	app       fyne.App
	window    fyne.Window
}

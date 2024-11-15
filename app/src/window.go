package main

import (
	"fyne.io/fyne/v2"
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

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type ProcWindow struct {
	x                float32
	y                float32
	width            float32
	height           float32
	app              fyne.App
	window           fyne.Window
	data             binding.ExternalStringList
	background_color color.NRGBA
	window_contents  *fyne.Container
}

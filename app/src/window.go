package main

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type ProcWindow struct {
	x                    float32
	y                    float32
	width                float32
	height               float32
	dark_mode            bool
	is_init              bool
	app                  fyne.App
	window               fyne.Window
	data                 []string // holds the list data
	background_color     color.NRGBA
	background_container *fyne.Container
	window_contents      *fyne.Container
}

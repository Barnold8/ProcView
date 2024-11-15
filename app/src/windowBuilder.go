package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type IWindowBuilder interface {
	InitialiseWindow() IWindowBuilder
	SetWindowPosition(x float32, y float32) IWindowBuilder
	SetWindowTitle(title string) IWindowBuilder
	SetWindowContainer(container *fyne.Container) IWindowBuilder
	SetWindowSize(width float32, height float32) IWindowBuilder
	Build() ProcWindow
}

type ConcreteWindowBuilder struct {
	processWindow ProcWindow
}

func (w *ConcreteWindowBuilder) InitialiseWindow() IWindowBuilder {
	w.processWindow.app = app.New()
	w.processWindow.window = w.processWindow.app.NewWindow("")
	return w
}

func (w *ConcreteWindowBuilder) CheckInit() {
	if w.processWindow.app == nil {
		log.Fatal("Window is not initialised in memory, run InitialiseWindow() first")
	}
}

func (w *ConcreteWindowBuilder) SetWindowPosition(x float32, y float32) IWindowBuilder {
	w.CheckInit()
	w.processWindow.x = x
	w.processWindow.window.Content().Position().AddXY(x, y)
	return w
}

func (w *ConcreteWindowBuilder) SetWindowSize(width float32, height float32) IWindowBuilder {
	w.CheckInit()

	w.processWindow.width = width
	w.processWindow.height = height

	return w
}

func (w *ConcreteWindowBuilder) SetWindowTitle(title string) IWindowBuilder {
	w.CheckInit()
	w.processWindow.window.SetTitle(title)
	return w
}

func (w *ConcreteWindowBuilder) SetWindowContainer(container *fyne.Container) IWindowBuilder { // All of windows contents must be passed in here to have a functioning window!
	w.CheckInit()
	w.processWindow.window_contents = container
	return w
}

func (w *ConcreteWindowBuilder) Build() ProcWindow {
	w.CheckInit()

	dimension_pair := ValidateDimensions(w.processWindow.width, w.processWindow.height)

	w.processWindow.window.Resize(fyne.NewSize(dimension_pair.a, dimension_pair.b))

	if w.processWindow.window_contents != nil {
		w.processWindow.window.SetContent(w.processWindow.window_contents)
	}

	return w.processWindow
}

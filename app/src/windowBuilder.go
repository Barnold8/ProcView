package main

import (
	"log"

	"fyne.io/fyne/v2/app"
)

type IWindowBuilder interface {
	InitialiseWindow() IWindowBuilder
	SetWindowPosition(x float32, y float32) IWindowBuilder
	SetWindowTitle(title string) IWindowBuilder
	Build() ProcWindow
}

type ConcreteWindowBuilder struct {
	processWindow ProcWindow
}

func (w *ConcreteWindowBuilder) InitialiseWindow() IWindowBuilder {
	w.processWindow.app = app.New()
	w.processWindow.window = w.processWindow.app.NewWindow("")
	w.processWindow.is_init = true
	return w
}

func (w *ConcreteWindowBuilder) CheckInit() {
	if w.processWindow.is_init == false {
		log.Fatal("Window is not initialised in memory, run InitialiseWindow() first")
	}
}

func (w *ConcreteWindowBuilder) SetWindowPosition(x float32, y float32) IWindowBuilder {
	w.CheckInit()
	w.processWindow.x = x
	w.processWindow.window.Content().Position().AddXY(x, y)
	return w
}

func (w *ConcreteWindowBuilder) SetWindowTitle(title string) IWindowBuilder {
	w.CheckInit()
	w.processWindow.window.SetTitle(title)
	return w
}

func (w *ConcreteWindowBuilder) Build() ProcWindow {
	w.CheckInit()
	return w.processWindow
}

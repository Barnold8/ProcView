package main

type IWindowBuilder interface {
	SetWindowPosition(x float32, y float32) IWindowBuilder
	SetWindowTitle(title string) IWindowBuilder
}

type ConcreteWindowBuilder struct {
	processWindow ProcWindow
}

func (w *ConcreteWindowBuilder) SetWindowPosition(x float32, y float32) IWindowBuilder {
	w.processWindow.x = x

	w.processWindow.window.Content().Position().AddXY(x, y)
	return w
}

func (w *ConcreteWindowBuilder) SetWindowTitle(title string) IWindowBuilder {
	w.processWindow.window.SetTitle(title)
	return w
}

func (w *ConcreteWindowBuilder) Build() ProcWindow { return w.processWindow }

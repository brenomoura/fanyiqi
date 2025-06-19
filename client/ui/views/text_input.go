package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CustomInput struct {
	widget.Entry
	fyne.Tabbable
	window *fyne.Window
}

func NewInput(window *fyne.Window) *CustomInput {
	entry := &CustomInput{}
	entry.ExtendBaseWidget(entry)
	entry.MultiLine = true
	entry.Wrapping = fyne.TextWrapBreak
	entry.setWindow(window)
	entry.SetPlaceHolder("Type here what do you want to translate")

	return entry
}

func (e *CustomInput) setWindow(window *fyne.Window) {
	e.window = window
}

func (e *CustomInput) onEscape() {
	if e.window != nil {
		(*e.window).Close()
	}
}

func (e *CustomInput) AcceptsTab() bool {
	return false
}

func (e *CustomInput) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEscape:
		e.onEscape()
	default:
		e.Entry.KeyDown(key)
	}
}

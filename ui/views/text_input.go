package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type customInput struct {
	widget.Entry
	window *fyne.Window
}

func NewInput(window *fyne.Window) *customInput {
	entry := &customInput{}
	entry.ExtendBaseWidget(entry)
	entry.MultiLine = true
	entry.Wrapping = fyne.TextWrapBreak
	entry.setWindow(window)
	entry.SetPlaceHolder("Type here what do you want to translate")
	return entry
}

func (e *customInput) setWindow(window *fyne.Window) {
	e.window = window
}

func (e *customInput) onEscape() {
	if e.window != nil {
		(*e.window).Close()
	}
}

func (e *customInput) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEscape:
		e.onEscape()
	default:
		e.Entry.KeyDown(key)
	}
}

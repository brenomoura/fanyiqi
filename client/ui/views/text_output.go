package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CustomOutput struct {
	widget.Entry
	window *fyne.Window
}

func NewOutput(window *fyne.Window) *CustomOutput {
	output := &CustomOutput{}
	output.ExtendBaseWidget(output)
	output.MultiLine = true
	output.setWindow(window)
	output.Wrapping = fyne.TextWrapBreak
	output.Disable()

	return output
}

func (e *CustomOutput) setWindow(window *fyne.Window) {
	e.window = window
}

func (e *CustomOutput) onEscape() {
	if e.window != nil {
		(*e.window).Close()
	}
}

func (e *CustomOutput) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEscape:
		e.onEscape()
	}
}

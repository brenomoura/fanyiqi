package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type customOutput struct {
	widget.Entry
	window *fyne.Window
}

func NewOutput(window *fyne.Window) *customOutput {
	output := &customOutput{}
	output.ExtendBaseWidget(output)
	output.MultiLine = true
	output.setWindow(window)
	output.Wrapping = fyne.TextWrapBreak
	output.Disable()
	return output
}

func (e *customOutput) setWindow(window *fyne.Window) {
	e.window = window
}

func (e *customOutput) onEscape() {
	if e.window != nil {
		(*e.window).Close()
	}
}

func (e *customOutput) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEscape:
		e.onEscape()
	}
}

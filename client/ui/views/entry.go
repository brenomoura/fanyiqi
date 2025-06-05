package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type customEntry struct {
	widget.Entry
	fyne.Tabbable
	window *fyne.Window
}

func NewCustomEntry(window *fyne.Window, placeholderText string, multiline bool) *customInput {
	entry := &customInput{}
	entry.ExtendBaseWidget(entry)
	entry.MultiLine = multiline
	entry.Wrapping = fyne.TextWrapBreak
	entry.setWindow(window)
	entry.SetPlaceHolder(placeholderText)

	return entry
}

func (e *customEntry) setWindow(window *fyne.Window) {
	e.window = window
}

func (e *customEntry) onEscape() {
	if e.window != nil {
		(*e.window).Close()
	}
}

func (e *customEntry) AcceptsTab() bool {
	return false
}

func (e *customEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEscape:
		e.onEscape()
	default:
		e.Entry.KeyDown(key)
	}
}

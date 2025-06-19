package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CustomEntry struct {
	widget.Entry
	fyne.Tabbable
	window *fyne.Window
}

func NewCustomEntry(window *fyne.Window, placeholderText string, multiline bool) *CustomEntry {
	entry := &CustomEntry{}
	entry.ExtendBaseWidget(entry)
	entry.MultiLine = multiline
	entry.Wrapping = fyne.TextWrapBreak
	entry.setWindow(window)
	entry.SetPlaceHolder(placeholderText)

	return entry
}

func (e *CustomEntry) setWindow(window *fyne.Window) {
	e.window = window
}

func (e *CustomEntry) onEscape() {
	if e.window != nil {
		(*e.window).Close()
	}
}

func (e *CustomEntry) AcceptsTab() bool {
	return false
}

func (e *CustomEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEscape:
		e.onEscape()
	default:
		e.Entry.KeyDown(key)
	}
}

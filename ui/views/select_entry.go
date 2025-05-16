package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type customSelectEntry struct {
	widget.SelectEntry
	window *fyne.Window
}

type CustomSelectEntryParams struct {
	Window  *fyne.Window
	Options []string
}

func NewCustomSelectEntry(params CustomSelectEntryParams) *customSelectEntry {
	selectInput := &customSelectEntry{}
	selectInput.ExtendBaseWidget(selectInput)
	selectInput.SelectEntry.SetOptions(params.Options)
	selectInput.setWindow(params.Window)
	return selectInput
}

func (e *customSelectEntry) setWindow(window *fyne.Window) {
	e.window = window
}

func (e *customSelectEntry) onEscape() {
	if e.window != nil {
		(*e.window).Close()
	}
}

func (e *customSelectEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyEscape:
		e.onEscape()
	default:
		e.Entry.KeyDown(key)
	}
}

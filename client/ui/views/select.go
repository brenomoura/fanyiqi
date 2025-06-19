package views

import (
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CustomSelect struct {
	widget.Select
	window  *fyne.Window
	options []string
}

type CustomSelectEntryParams struct {
	Window  *fyne.Window
	Options []string
}

func NewCustomSelectEntry(params CustomSelectEntryParams) *CustomSelect {
	selectInput := &CustomSelect{}
	selectInput.ExtendBaseWidget(selectInput)
	selectInput.Select.SetOptions(params.Options)
	selectInput.setWindow(params.Window)
	selectInput.options = params.Options
	return selectInput
}

func (e *CustomSelect) setWindow(window *fyne.Window) {
	e.window = window
}

func (e *CustomSelect) onEscape() {
	if e.window != nil {
		(*e.window).Close()
	}
}

func (e *CustomSelect) TypedKey(event *fyne.KeyEvent) {
	switch event.Name {
	case fyne.KeyUp:
		selectedOption := e.Select.Selected
		if len(e.Select.Selected) == 0 {
			return
		}
		currentIndex := slices.Index(e.options, selectedOption)
		if currentIndex <= 0 {
			return
		}
		e.Select.SetSelected(e.options[currentIndex-1])
		e.Select.Refresh()
	case fyne.KeyDown:
		selectedOption := e.Select.Selected
		if len(selectedOption) == 0 {
			e.Select.SetSelected(e.options[0])
			return
		}
		currentIndex := slices.Index(e.options, selectedOption)
		if currentIndex < 0 || currentIndex >= len(e.options)-1 {
			return
		}
		e.Select.SetSelected(e.options[currentIndex+1])
		e.Select.Refresh()
	case fyne.KeyEscape:
		e.onEscape()
	default:
		e.Select.TypedKey(event)
	}
}

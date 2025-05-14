package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct {
	fyne.Theme
}

func (t *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameDisabled:
		// the disabled text color is not too good to read, due to this, I just set it return as nil.
		// and also, it easier to create a custom theme than modify the entry to have it the way I'd like.
		return nil
	default:
		return t.Theme.Color(name, variant)
	}
}

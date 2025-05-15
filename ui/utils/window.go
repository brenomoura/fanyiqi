package utils

import (
	"fyne.io/fyne/v2"
	"github.com/kbinani/screenshot"
)

func SetWindowSize() fyne.Size {
	if screenshot.NumActiveDisplays() > 0 {
		// #0 is the main monitor
		bounds := screenshot.GetDisplayBounds(0)
		return fyne.NewSize(float32(bounds.Dx()/3), float32(bounds.Dy())/3)
	}
	return fyne.NewSize(800, 800)
}

func Close(window fyne.Window) {
	window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		println(event)
		if event.Name == fyne.KeyEscape {
			window.Close()
		}
	})
}

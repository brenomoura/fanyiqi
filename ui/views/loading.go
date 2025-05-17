package views

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/theme"
)

type Loading struct {
	Rectangle      *canvas.Rectangle
	Text           *canvas.Text
	Container      *fyne.Container
	LoadingMessage string
	loading        bool
}

func NewLoading() *Loading {
	rec := canvas.NewRectangle(color.Gray16{Y: 0x2FFF})
	rec.Hide()
	text := canvas.NewText("", theme.ForegroundColor())
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{}
	cont := container.New(layout.NewStackLayout(), rec, text)
	return &Loading{
		Rectangle:      rec,
		Text:           text,
		Container:      cont,
		LoadingMessage: "Carregando meu chapa",
		loading:        false,
	}
}

func (l *Loading) animate() {
	go func() {
		dots := []string{".  ", ".. ", "..."}
		index := 0
		for l.loading {
			time.Sleep(time.Millisecond * 400)
			fyne.Do(func() {
				l.Text.Text = l.LoadingMessage + dots[index%len(dots)]
				canvas.Refresh(l.Text)
				canvas.Refresh(l.Rectangle)
			})
			index++
		}
	}()
}

func (l *Loading) SetLoading(isLoading bool) {
	l.loading = isLoading
	if isLoading {
		l.Text.Show()
		l.Rectangle.Show()
		l.animate()
	} else {
		l.Text.Hide()
		l.Rectangle.Hide()
	}
}

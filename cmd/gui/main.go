package main

import (
	"github.com/brenomoura/fanyiqi/ui/utils"
	"github.com/brenomoura/fanyiqi/ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("fanyiqi")
	app.Settings().SetTheme(&views.CustomTheme{Theme: theme.DefaultTheme()})
	window.Resize(utils.SetWindowSize())
	window.CenterOnScreen()
	window.SetContent(makeUI(window))
	utils.Close(window)
	window.ShowAndRun()
}

func makeUI(w fyne.Window) fyne.CanvasObject {
	header := canvas.NewText("fanyiqi", theme.PrimaryColor())
	header.TextSize = 25
	input := views.NewInput(&w)

	output := views.NewOutput(&w)

	input.OnChanged = func(typedChar string) {
		output.Text = typedChar
		output.Refresh()
	}

	clear := widget.NewButtonWithIcon("clear", theme.ContentClearIcon(), func() {
		output.Text = ""
		output.Refresh()
		input.Text = ""
		input.Refresh()
	})

	content := container.NewGridWithColumns(2, input, output)

	return container.NewBorder(
		header,
		container.NewHBox(clear),
		nil,
		nil,
		content,
	)
}

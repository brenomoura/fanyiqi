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
	header := canvas.NewText("fanyiqi", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 25

	footer := canvas.NewText("Press BATATA to see the shorcuts", theme.Color(theme.ColorNamePrimary))
	footer.TextSize = 10

	input := views.NewInput(&w)
	output := views.NewOutput(&w)

	input.OnChanged = func(typedChar string) {
		output.Text = typedChar
		output.Refresh()
	}

	inputSelectEntry := widget.NewSelectEntry([]string{"Português", "Inglês"})
	inputSelectEntry.OnChanged = func(typedChar string) {}
	outputSelectEntry := widget.NewSelectEntry([]string{"Português", "Inglês"})

	inputView := container.NewBorder(inputSelectEntry, nil, nil, nil, input)
	outputView := container.NewBorder(outputSelectEntry, nil, nil, nil, output)

	content := container.NewGridWithColumns(
		2,
		inputView,
		outputView,
	)

	return container.NewBorder(
		header,
		footer,
		nil,
		nil,
		content,
	)
}

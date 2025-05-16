package main

import (
	"time"

	"github.com/brenomoura/fanyiqi/ui/utils"
	"github.com/brenomoura/fanyiqi/ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	bar := widget.NewProgressBarInfinite()
	bar.Theme().Color(theme.ColorNameDisabled, theme.VariantDark)
	bar.Hide()

	outputStack := container.New(layout.NewStackLayout(), output, bar)

	clearButton := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), func() {
		input.Text = ""
		input.Refresh()
		output.Text = ""
		output.Refresh()
	})

	input.OnChanged = func(typedChar string) {
		output.Text = ""
		output.Refresh()
		bar.Show()
		go func() {
			time.Sleep(time.Millisecond * 1000)
			fyne.Do(func() {
				bar.Hide()
				output.Text = typedChar
				output.Refresh()
			})
		}()
	}

	inputSelectEntry := views.NewCustomSelectEntry(views.CustomSelectEntryParams{
		Window:  &w,
		Options: []string{"Português", "Inglês"},
	})
	outputSelectEntry := views.NewCustomSelectEntry(views.CustomSelectEntryParams{
		Window:  &w,
		Options: []string{"Português", "Inglês"},
	})

	inputView := container.NewBorder(inputSelectEntry, clearButton, nil, nil, input)
	outputView := container.NewBorder(outputSelectEntry, nil, nil, nil, outputStack)

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

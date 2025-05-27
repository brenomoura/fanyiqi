package main

import (
	// "os"
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
	"golang.design/x/clipboard"
)

func main() {
	// println(os.UserConfigDir())
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	app := app.New()
	window := app.NewWindow("fanyiqi")
	app.Settings().SetTheme(&views.CustomTheme{Theme: theme.DefaultTheme()})
	window.Resize(utils.SetWindowSize())
	window.CenterOnScreen()
	header := canvas.NewText("fanyiqi", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 25

	footer := canvas.NewText("Press BATATA to setup a translator provider", theme.Color(theme.ColorNamePrimary))
	footer.TextSize = 10

	input := views.NewInput(&window)
	output := views.NewOutput(&window)

	loading := views.NewLoading()
	outputStack := container.New(layout.NewStackLayout(), output, loading.Container)

	clearButton := widget.NewButtonWithIcon("Clear (Ctrl + PQP)", theme.ContentClearIcon(), func() {
		input.Text = ""
		input.Refresh()
		output.Text = ""
		output.Refresh()
	})

	settingsButton := widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {})

	buttons := container.NewGridWithColumns(2, clearButton, settingsButton)

	input.OnChanged = func(typedChar string) {
		output.Text = ""
		output.Refresh()
		loading.SetLoading(true)
		go func() {
			time.Sleep(time.Millisecond * 1000)
			loading.SetLoading(false)
			fyne.Do(func() {
				output.Text = typedChar
				output.Refresh()
			})
		}()
	}

	clipboardBytes := clipboard.Read(clipboard.FmtText)
	if clipboardBytes != nil {
		clipboardText := string(clipboard.Read(clipboard.FmtText))
		if len(clipboardText) > 0 {
			input.Text = clipboardText
			input.OnChanged(input.Text)
		}
	}

	options := []string{"Português", "Inglês"}

	inputSelectEntry := views.NewCustomSelectEntry(views.CustomSelectEntryParams{
		Window:  &window,
		Options: options,
	})

	outputSelectEntry := views.NewCustomSelectEntry(views.CustomSelectEntryParams{
		Window:  &window,
		Options: options,
	})

	inputView := container.NewBorder(inputSelectEntry, buttons, nil, nil, input)
	outputView := container.NewBorder(outputSelectEntry, nil, nil, nil, outputStack)

	mainContent := container.NewGridWithColumns(
		2,
		inputView,
		outputView,
	)

	label1 := widget.NewLabel("Label 1")
	value1 := widget.NewLabel("Value")
	label2 := widget.NewLabel("Label 2")
	value2 := widget.NewLabel("Something")
	settingsContent := container.New(
		layout.NewFormLayout(),
		label1,
		value1,
		label2,
		value2,
	)

	ui := container.NewBorder(
		header,
		footer,
		nil,
		nil,
		mainContent,
	)

	settingsButton.OnTapped = func() {
		ui.Objects[0] = settingsContent
		println("test")
		ui.Refresh()
	}

	window.SetContent(ui)
	window.Canvas().Focus(input)
	utils.Close(window)
	window.ShowAndRun()
}

package main

import (
	"time"

	"github.com/brenomoura/fanyiqi/internal/config"
	"github.com/brenomoura/fanyiqi/internal/translator"
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

	input := views.NewInput(&window)
	output := views.NewOutput(&window)

	loading := views.NewLoading()
	outputStack := container.New(layout.NewStackLayout(), output, loading.Container)

	clearButton := widget.NewButtonWithIcon("Clear input text (Ctrl + PQP)", theme.ContentClearIcon(), func() {
		input.Text = ""
		input.Refresh()
		output.Text = ""
		output.Refresh()
	})

	settingsButton := widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), func() {})

	translationButton := widget.NewButtonWithIcon("Translate!", theme.DocumentIcon(), func() {})

	buttons := container.NewGridWithColumns(2, clearButton, settingsButton)

	appConfig, err := config.LoadEncryptedConfig()
	if err != nil {
		window.Close()
	}

	println("Loaded config:", appConfig.APIKey, *appConfig.SourceLanguage, *appConfig.TargetLanguage)

	apiKeyEntry := views.NewCustomEntry(
		&window,
		"Insert here the API key from provider...",
		true,
	)

	saveAPIKeyButton := widget.NewButtonWithIcon("Save API Key", theme.DocumentSaveIcon(), func() {})
	settingsButtons := container.NewGridWithColumns(2, translationButton, saveAPIKeyButton)
	settingsContent := container.New(layout.NewGridLayoutWithRows(2), apiKeyEntry, container.New(layout.NewCenterLayout(), settingsButtons))

	ui := container.NewBorder(
		header,
		nil,
		nil,
		nil,
		settingsContent,
	)

	saveAPIKeyButton.OnTapped = func() {
		apiKey := apiKeyEntry.Text
		if len(apiKey) == 0 {
			apiKeyEntry.SetPlaceHolder("Please type your API Key before saving it.")
			apiKeyEntry.Refresh()
			return
		}
		appConfig.APIKey = apiKey
		config.SaveEncryptedConfig(*appConfig)
		apiKeyEntry.Text = ""
		apiKeyEntry.SetPlaceHolder("API Key saved successfully! You can change it anytime by typing the new key and click 'Save API Key'.")
		apiKeyEntry.Refresh()
	}

	if appConfig.APIKey == "" {
		window.SetContent(ui)
		window.Canvas().Focus(input)
		utils.Close(window)
		window.ShowAndRun()
		return
	}

	if appConfig != nil && appConfig.APIKey != "" {
		apiKeyEntry.SetPlaceHolder("API Key already set. To overwrite, type the new key and click 'Save API Key'.")
	}

	translatorService := translator.NewTranslatorService(
		"http://localhost:8000/api/v1",
		appConfig.APIKey,
	)

	input.OnChanged = func(typedChar string) {
		output.Text = ""
		output.Refresh()
		loading.SetLoading(true)
		go func() {
			time.Sleep(time.Millisecond * 1000)
			loading.SetLoading(false)
			fyne.Do(func() {
				result, err := translatorService.Translate(translator.TranslationParams{
					Text:           typedChar,
					SourceLanguage: *appConfig.SourceLanguage,
					TargetLanguage: *appConfig.TargetLanguage,
				})
				if err != nil {
					// add some dialog to show the error
					// add logs
					println("Error translating text:", err)
					return
				}
				if result.TranslatedText == "" {
					println("Received empty translation result")
					return
				}
				output.Text = result.TranslatedText
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

	langsOptions, err := translatorService.GetLanguages()
	if err != nil {
		// add some dialog to show the error
		// add logs
		println("Error fetching languages:", err)
	}

	options := make([]string, 0, len(langsOptions))
	for _, lang := range langsOptions {
		options = append(options, lang[0])
	}

	inputSelectEntry := views.NewCustomSelectEntry(views.CustomSelectEntryParams{
		Window:  &window,
		Options: options,
	})

	if appConfig != nil && appConfig.SourceLanguage != nil {
		for _, lang := range langsOptions {
			if lang[1] == *appConfig.SourceLanguage {
				inputSelectEntry.SetSelected(lang[0])
				inputSelectEntry.Refresh()
				break
			}
		}
	}

	inputSelectEntry.OnChanged = func(selectedOption string) {
		if selectedOption == "" {
			return
		}
		var selectedLang string
		for _, lang := range langsOptions {
			if lang[0] == selectedOption {
				selectedLang = lang[1]
				break
			}
		}
		appConfig.SourceLanguage = &selectedLang
		config.SaveEncryptedConfig(*appConfig)
		inputSelectEntry.Refresh()
	}

	outputSelectEntry := views.NewCustomSelectEntry(views.CustomSelectEntryParams{
		Window:  &window,
		Options: options,
	})

	if appConfig != nil && appConfig.TargetLanguage != nil {
		for _, lang := range langsOptions {
			if lang[1] == *appConfig.TargetLanguage {
				outputSelectEntry.SetSelected(lang[0])
				outputSelectEntry.Refresh()
				break
			}
		}
	}

	outputSelectEntry.OnChanged = func(selectedOption string) {
		if selectedOption == "" {
			return
		}
		var selectedLang string
		for _, lang := range langsOptions {
			if lang[0] == selectedOption {
				selectedLang = lang[1]
				break
			}
		}
		appConfig.TargetLanguage = &selectedLang
		config.SaveEncryptedConfig(*appConfig)
		outputSelectEntry.Refresh()
	}

	inputView := container.NewBorder(inputSelectEntry, buttons, nil, nil, input)
	outputView := container.NewBorder(outputSelectEntry, nil, nil, nil, outputStack)

	mainContent := container.NewGridWithColumns(
		2,
		inputView,
		outputView,
	)

	if appConfig.APIKey == "" {
		ui.Objects[0] = mainContent
	}

	settingsButton.OnTapped = func() {
		ui.Objects[0] = settingsContent
		ui.Refresh()
	}

	translationButton.OnTapped = func() {
		ui.Objects[0] = mainContent
		ui.Refresh()
	}

	window.SetContent(ui)
	window.Canvas().Focus(input)
	utils.Close(window)
	window.ShowAndRun()
}

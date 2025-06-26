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

type Application struct {
	app               fyne.App
	window            fyne.Window
	config            *config.Config
	translator        *translator.TranslatorService
	ui                *fyne.Container
	mainContent       *fyne.Container
	settingsContent   *fyne.Container
	input             *views.CustomInput
	output            *views.CustomOutput
	loading           *views.Loading
	apiURLEntry       *views.CustomEntry
	apiKeyEntry       *views.CustomEntry
	inputSelectEntry  *views.CustomSelect
	outputSelectEntry *views.CustomSelect
	inputView         fyne.CanvasObject
	outputView        fyne.CanvasObject
	languageOptions   [][]string
	debounceTimer     *time.Timer
}

func main() {
	app := &Application{}
	app.initialize()
	app.run()
}

func (a *Application) initialize() {
	a.initClipboard()
	a.createAppWindow()
	a.loadConfig()
	a.setupUIComponents()
}

func (a *Application) run() {
	a.window.Canvas().Focus(a.input)
	utils.Close(a.window)
	a.window.ShowAndRun()
}

func (a *Application) initClipboard() {
	if err := clipboard.Init(); err != nil {
		panic(err)
	}
}

func (a *Application) createAppWindow() {
	a.app = app.New()
	a.window = a.app.NewWindow("fanyiqi")
	a.app.Settings().SetTheme(&views.CustomTheme{Theme: theme.DefaultTheme()})
	a.window.Resize(utils.SetWindowSize())
	a.window.CenterOnScreen()
}

func (a *Application) loadConfig() {
	cfg, err := config.LoadEncryptedConfig()
	if err != nil {
		a.window.Close()
	}
	a.config = cfg
}

func (a *Application) setupUIComponents() {
	a.setupHeader()
	a.setupSettingsUI()
	a.setupTranslatorService()
	a.setupLanguageSelection()
	a.setupMainUI()
	a.setupEventHandlers()
	a.setupClipboardIntegration()
	a.setupUI()
}

func (a *Application) setupHeader() *canvas.Text {
	header := canvas.NewText("fanyiqi", theme.Color(theme.ColorNamePrimary))
	header.TextSize = 25
	return header
}

func (a *Application) setupMainUI() {
	a.input = views.NewInput(&a.window)
	a.output = views.NewOutput(&a.window)
	a.loading = views.NewLoading()
	outputStack := container.New(layout.NewStackLayout(), a.output, a.loading.Container)

	clearButton := widget.NewButtonWithIcon("Clear", theme.ContentClearIcon(), a.handleClear)
	settingsButton := widget.NewButtonWithIcon("Settings", theme.SettingsIcon(), a.handleSettingsButton)
	swapLanguagesButton := widget.NewButtonWithIcon("Swap Languages", theme.ViewRefreshIcon(), a.handleSwapLanguagesButton)
	buttons := container.NewGridWithColumns(3, clearButton, settingsButton, swapLanguagesButton)

	inputSelect := a.inputSelectEntry
	outputSelect := a.outputSelectEntry

	a.inputView = container.NewBorder(inputSelect, buttons, nil, nil, a.input)
	a.outputView = container.NewBorder(outputSelect, nil, nil, nil, outputStack)

	a.mainContent = container.NewGridWithColumns(2, a.inputView, a.outputView)
}

func (a *Application) setupSettingsUI() {
	a.apiURLEntry = views.NewCustomEntry(
		&a.window,
		"Insert here the API URL from the provider...",
		false,
	)

	if a.config != nil && a.config.APIKey != "" {
		a.apiURLEntry.SetPlaceHolder("API URL already set. To overwrite, type the new URL and click 'Save'.")
	}
	a.apiKeyEntry = views.NewCustomEntry(
		&a.window,
		"Insert here the API key from the provider...",
		true,
	)

	if a.config != nil && a.config.APIKey != "" {
		a.apiKeyEntry.SetPlaceHolder("API Key already set. To overwrite, type the new key and click 'Save'.")
	}

	translationButton := widget.NewButtonWithIcon("Translate!", theme.DocumentIcon(), a.handleReturnButton)
	saveConfigsButton := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), a.saveConfigs)
	settingsButtons := container.NewGridWithColumns(2, translationButton, saveConfigsButton)
	a.settingsContent = container.New(
		layout.NewVBoxLayout(),
		a.apiURLEntry,
		a.apiKeyEntry,
		container.New(layout.NewCenterLayout(), settingsButtons),
	)
}

func (a *Application) setupUI() {
	a.ui = container.NewBorder(
		a.setupHeader(),
		nil,
		nil,
		nil,
		a.settingsContent,
	)

	if a.config.APIKey != "" {
		a.ui.Objects[0] = a.mainContent
	}
	a.window.SetContent(a.ui)
}

func (a *Application) setupEventHandlers() {
	a.input.OnChanged = a.handleInputChanged
}

func (a *Application) setupTranslatorService() {
	a.translator = translator.NewTranslatorService(
		a.config.APIURL,
		a.config.APIKey,
	)
}

func (a *Application) getLanguages() ([][]string, error) {
	if a.config.APIKey == "" {
		return [][]string{}, nil
	}
	if len(a.languageOptions) > 0 {
		return a.languageOptions, nil
	}

	languages, err := a.translator.GetLanguages()
	if err != nil {
		println("Error fetching languages:", err)
		return nil, err
	}

	return languages, nil
}

func (a *Application) getSelectOptions() []string {
	if len(a.languageOptions) == 0 {
		return nil
	}

	options := make([]string, 0, len(a.languageOptions))
	for _, lang := range a.languageOptions {
		options = append(options, lang[0])
	}
	return options
}

func (a *Application) setupLanguageSelection() {
	var options []string
	langsOptions, err := a.getLanguages()
	if err != nil {
		println("Error setting up language selection:", err)
		return
	}
	a.languageOptions = langsOptions
	options = a.getSelectOptions()
	a.setupInputSelect(options, a.languageOptions)
	a.setupOutputSelect(options, a.languageOptions)
}

func (a *Application) setupInputSelect(options []string, langsOptions [][]string) {
	a.inputSelectEntry = views.NewCustomSelectEntry(views.CustomSelectEntryParams{
		Window:  &a.window,
		Options: options,
	})

	if a.config.SourceLanguage != nil {
		a.setSelectedLanguage(a.inputSelectEntry, langsOptions, *a.config.SourceLanguage)
	}

	a.inputSelectEntry.OnChanged = a.createLanguageChangeHandler("SourceLanguage", a.config)
}

func (a *Application) setupOutputSelect(options []string, langsOptions [][]string) {
	a.outputSelectEntry = views.NewCustomSelectEntry(views.CustomSelectEntryParams{
		Window:  &a.window,
		Options: options,
	})

	if a.config.TargetLanguage != nil {
		a.setSelectedLanguage(a.outputSelectEntry, langsOptions, *a.config.TargetLanguage)
	}

	a.outputSelectEntry.OnChanged = a.createLanguageChangeHandler("TargetLanguage", a.config)
}

func (a *Application) setSelectedLanguage(selectEntry *views.CustomSelect, langsOptions [][]string, langCode string) {
	for _, lang := range langsOptions {
		if lang[1] == langCode {
			selectEntry.SetSelected(lang[0])
			selectEntry.Refresh()
			break
		}
	}
}

func (a *Application) createLanguageChangeHandler(field string, configStruct *config.Config) func(string) {
	return func(selectedOption string) {
		if selectedOption == "" {
			return
		}

		selectedLang := ""
		for _, lang := range a.languageOptions {
			if lang[0] == selectedOption {
				selectedLang = lang[1]
				break
			}
		}
		switch field {
		case "SourceLanguage":
			configStruct.SourceLanguage = &selectedLang
		case "TargetLanguage":
			configStruct.TargetLanguage = &selectedLang
		}
		config.SaveEncryptedConfig(*configStruct)
	}
}

func (a *Application) setupClipboardIntegration() {
	clipboardBytes := clipboard.Read(clipboard.FmtText)
	if clipboardBytes == nil {
		return
	}

	clipboardText := string(clipboardBytes)
	if len(clipboardText) > 0 {
		a.input.Text = clipboardText
		a.input.OnChanged(a.input.Text)
	}
}

func (a *Application) handleClear() {
	a.input.Text = ""
	a.input.Refresh()
	a.output.Text = ""
	a.output.Refresh()
}

func (a *Application) handleSettingsButton() {
	a.ui.Objects[0] = a.settingsContent
	a.ui.Refresh()
}

func (a *Application) handleSwapLanguagesButton() {
	if a.config.SourceLanguage == nil || a.config.TargetLanguage == nil {
		println("Source or target language is not set.")
		return
	}

	a.config.SourceLanguage, a.config.TargetLanguage = a.config.TargetLanguage, a.config.SourceLanguage

	config.SaveEncryptedConfig(*a.config)

	selectedInput, selectedOutput := "", ""
	for _, lang := range a.languageOptions {
		if lang[1] == *a.config.SourceLanguage {
			selectedInput = lang[0]
		}
		if lang[1] == *a.config.TargetLanguage {
			selectedOutput = lang[0]
		}
	}

	a.inputSelectEntry.SetSelected(selectedInput)
	a.outputSelectEntry.SetSelected(selectedOutput)

	a.inputSelectEntry.Refresh()
	a.outputSelectEntry.Refresh()
	a.input.OnChanged(a.input.Text)
}

func (a *Application) handleReturnButton() {
	a.ui.Objects[0] = a.mainContent
	a.ui.Refresh()
	if len(a.languageOptions) == 0 {
		a.setupTranslatorService()
		langs, err := a.getLanguages()
		if err != nil {
			println("Error fetching languages:", err)
			return
		}
		a.languageOptions = langs
		options := a.getSelectOptions()
		a.inputSelectEntry.SetOptions(options)
		a.outputSelectEntry.SetOptions(options)
		a.inputSelectEntry.Refresh()
		a.outputSelectEntry.Refresh()
	}
}

func (a *Application) saveConfigs() {
	apiURL := a.apiURLEntry.Text
	if apiURL == "" {
		a.apiURLEntry.SetPlaceHolder("Please type your API URL before saving it.")
		a.apiURLEntry.Refresh()
		return
	}

	apiKey := a.apiKeyEntry.Text
	if apiKey == "" {
		a.apiKeyEntry.SetPlaceHolder("Please type your API Key before saving it.")
		a.apiKeyEntry.Refresh()
		return
	}

	a.config.APIURL = apiURL
	a.config.APIKey = apiKey
	config.SaveEncryptedConfig(*a.config)

	a.apiURLEntry.Text = ""
	a.apiKeyEntry.Text = ""
	a.apiURLEntry.SetPlaceHolder("API URL saved successfully! You can change it anytime by typing the new URL and click 'Save'.")
	a.apiKeyEntry.SetPlaceHolder("API Key saved successfully! You can change it anytime by typing the new key and click 'Save'.")
	a.apiURLEntry.Refresh()
	a.apiKeyEntry.Refresh()
}

func (a *Application) handleInputChanged(typedChar string) {
	if typedChar == "" || a.config.SourceLanguage == nil || a.config.TargetLanguage == nil {
		a.output.Text = ""
		a.output.Refresh()
		return
	}

	if a.debounceTimer != nil {
		a.debounceTimer.Stop()
	}

	a.output.Text = ""
	a.output.SetPlaceHolder("")
	a.output.Refresh()
	a.debounceTimer = time.AfterFunc(500*time.Millisecond, func() {
		a.loading.SetLoading(true)
		result, err := a.translator.Translate(translator.TranslationParams{
			Text:           typedChar,
			SourceLanguage: *a.config.SourceLanguage,
			TargetLanguage: *a.config.TargetLanguage,
		})

		if err != nil {
			a.output.Text = ""
			a.output.SetPlaceHolder("Error translating text")
			a.loading.SetLoading(false)
			return
		}

		a.output.Text = result.TranslatedText
		fyne.Do(func() {
			a.output.Refresh()
			a.loading.SetLoading(false)
		})
	})

}

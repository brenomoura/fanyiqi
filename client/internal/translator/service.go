package translator

import (
	"github.com/brenomoura/fanyiqi/pkg/http"
)

type TranslationParams struct {
	Text           string `json:"text"`
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
	Model          string `json:"model,omitempty"`
}

type TranslationResult struct {
	TranslatedText string `json:"translated_text"`
}

type TranslatorService struct {
	Client *http.HTTPClient
}

func NewTranslatorService(baseURL, token string) *TranslatorService {
	return &TranslatorService{
		Client: http.NewHTTPClient(baseURL, token),
	}
}

func (t *TranslatorService) Translate(params TranslationParams) (TranslationResult, error) {
	var res TranslationResult
	err := t.Client.PostJSON("/translate", params, &res)
	return res, err
}

func (t *TranslatorService) GetLanguages() ([][]string, error) {
	var res [][]string
	err := t.Client.GetJSON("/languages", &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

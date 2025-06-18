package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	BaseURL   string
	AuthToken string
	Client    *http.Client
}

func NewHTTPClient(baseURL, token string) *HTTPClient {
	return &HTTPClient{
		BaseURL:   baseURL,
		AuthToken: token,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *HTTPClient) PostJSON(path string, body any, result any) error {
	fullURL := h.BaseURL + path

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if h.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+h.AuthToken)
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		raw, _ := io.ReadAll(resp.Body)
		return errors.New(string(raw))
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func (h *HTTPClient) GetJSON(path string, result any) error {
	fullURL := h.BaseURL + path

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return err
	}

	if h.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+h.AuthToken)
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		raw, _ := io.ReadAll(resp.Body)
		return errors.New(string(raw))
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

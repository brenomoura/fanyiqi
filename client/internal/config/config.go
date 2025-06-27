package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type TokenData struct {
	Token string `json:"token"`
}

type Config struct {
	APIURL         string  `json:"api_url"`
	APIKey         string  `json:"api_key"`
	Provider       *string `json:"provider,omitempty"`
	SourceLanguage *string `json:"source_language,omitempty"`
	TargetLanguage *string `json:"target_language,omitempty"`
}

func getKeyFilePath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(userConfigDir, "fanyiqi")
	if err := os.MkdirAll(appDir, 0700); err != nil {
		return "", err
	}
	// TODO: Use a more secure place
	return filepath.Join(appDir, "key.bin"), nil
}

func loadOrCreateKey() ([]byte, error) {
	keyFile, err := getKeyFilePath()

	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(keyFile); err == nil {
		return os.ReadFile(keyFile)
	}

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	if err := os.WriteFile(keyFile, key, 0600); err != nil {
		return nil, err
	}
	return key, nil
}

func getOrCreateConfigFilePath() *string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return nil
	}

	appDir := filepath.Join(userConfigDir, "fanyiqi")

	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		if err := os.MkdirAll(appDir, 0700); err != nil {
			return nil
		}
	}

	tokenFilePath := filepath.Join(appDir, "config.enc")
	return &tokenFilePath
}

func SaveEncryptedConfig(config Config) error {
	key, err := loadOrCreateKey()
	if err != nil {
		return err
	}

	plainText, err := json.Marshal(config)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return err
	}

	cipherText := aesGCM.Seal(nonce, nonce, plainText, nil)
	tokenFilePath := getOrCreateConfigFilePath()
	if tokenFilePath == nil {
		return errors.New("failed to get config file path")
	}
	return os.WriteFile(*tokenFilePath, cipherText, 0600)
}

func LoadEncryptedConfig() (*Config, error) {
	configFilePath := getOrCreateConfigFilePath()
	if configFilePath == nil {
		return nil, errors.New("failed to get config file path")
	}

	cipherText, err := os.ReadFile(*configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			emptyConfig := &Config{APIKey: ""}
			SaveEncryptedConfig(*emptyConfig)
			cipherText, err = os.ReadFile(*configFilePath)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	key, err := loadOrCreateKey()

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("invalid config file")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	var data Config
	err = json.Unmarshal(plainText, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

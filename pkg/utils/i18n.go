package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type I18nMap map[string]map[string]string // field -> tag -> message

var ValidationMessages map[string]I18nMap

func LoadValidationMessages() error {
	ValidationMessages = make(map[string]I18nMap)

	files := []string{"pkg/i18n/validation.en.json", "pkg/i18n/validation.id.json"}
	for _, file := range files {
		lang := filepath.Base(file)[11:13] // ambil "en", "id", dll

		f, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to load i18n file: %w", err)
		}

		var data map[string]map[string]string
		if err := json.Unmarshal(f, &data); err != nil {
			return fmt.Errorf("failed to parse i18n JSON: %w", err)
		}

		ValidationMessages[lang] = data
	}
	return nil
}

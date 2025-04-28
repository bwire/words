package io

import (
	"encoding/json"
	"os"
	"words/internal/model"
)

func PersistResultToFile(path string, words []model.WordEntry) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(words)
}

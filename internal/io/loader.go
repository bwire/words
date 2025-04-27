package io

import (
	"encoding/json"
	"fmt"
	"os"
	"words/internal/model"
)

func ReadWordList(path string) ([]model.WordEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var words []model.WordEntry

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&words); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return words, nil
}
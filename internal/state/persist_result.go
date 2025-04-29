package state

import (
	"fmt"
	"words/internal/config"
	"words/internal/io"
	"words/internal/model"
)

type StatePersistResult struct {
	NewDictionary []model.WordEntry
}

func (s StatePersistResult) Execute(sm *StateMachine) (State, error) {
	fmt.Println(ProcessMessage("Persisting results..."))

	err := io.PersistResultToFile(config.Vocabulary, s.NewDictionary)
	if err != nil {
		return nil, fmt.Errorf("failed to save dictitionary: %w", err)
	}

	fmt.Println(ResultMessage("Dictionary successfully update! Exiting..."))
	return nil, nil
}

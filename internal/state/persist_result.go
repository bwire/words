package state

import (
	"fmt"
	"words/internal/config"
	"words/internal/io"
)

type StatePersistResult struct{}

func (s StatePersistResult) Execute(sm *StateMachine) (State, error) {
	fmt.Println("Persisting results...")

	err := io.PersistResult(config.Vocabulary, sm.ActiveWords)
	if err != nil {
		return nil, fmt.Errorf("failed to save dictitionary: %w", err)
	}

	fmt.Println("Vocabulary successfully saved!")
	return nil, nil
}

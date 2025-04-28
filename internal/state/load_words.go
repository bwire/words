package state

import (
	"fmt"
	"words/internal/config"
	"words/internal/io"
)

type StateLoadWords struct{}

func (s StateLoadWords) Execute(sm *StateMachine) (State, error) {
	fmt.Println("Loading active vocabulary...")

	words, err := io.ReadWordList(config.Vocabulary)
	if err != nil {
		return nil, err
	}

	sm.ActiveWords = words

	fmt.Printf("Active dictionary conatins %v words and phrases for learning!\n\n", len(sm.ActiveWords))
	return StateValidateDictionary{}, nil
}

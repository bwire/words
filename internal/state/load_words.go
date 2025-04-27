package state

import (
	"fmt"
	"words/internal/config"
	"words/internal/io"
)

type StateLoadWords struct {}

func (s StateLoadWords) Execute(sm *StateMachine) (State, error) {
	fmt.Println("Loading active vocabulary...")

	words, err := io.ReadWordList(config.Vocabulary)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	sm.ActiveWords = words
	return StateValidateDictionary{}, nil
}

package state

import (
	"fmt"
	"time"
	"words/internal/config"
)

type StateValidateDictionary struct {}

func (s StateValidateDictionary) Execute(sm *StateMachine) (State, error) {
	fmt.Println("Validating vocabulary...")

	if l := len(sm.ActiveWords); l < config.WordsInSet {
		fmt.Println("Not enough words in the dictionary! You should add some")
		return nil, fmt.Errorf("not enough words in the dictionary (%v)! You should add some", l)
	}
	for i := range sm.ActiveWords {
		word := &sm.ActiveWords[i]
		if word.StartDate == "" {
			word.StartDate = time.Now().Format("2006-01-02")
		}
		if word.HitsCount == 0 {
			word.HitsCount = word.Progress
		}
	}

	return StatePersistResult{}, nil

}

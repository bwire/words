package state

import (
	"fmt"
	"time"
	"words/internal/config"
)

type StateValidateDictionary struct{}

func (s StateValidateDictionary) Execute(sm *StateMachine) (State, error) {
	fmt.Println(ProcessMessage("Validating dictionary..."))

	if l := len(sm.ActiveWords); l < config.WordsInSet {
		return nil, fmt.Errorf("not enough words in the dictionary (%v)! You should add some", l)
	}

	for i := range sm.ActiveWords {
		word := &sm.ActiveWords[i]

		if word.Word == "" || word.Meaning == "" {
			return nil, fmt.Errorf("invalid word entry detected: %+v", word)
		}

		if word.StartDate == "" {
			word.StartDate = time.Now().Format(config.DateFormat)
		}
		if word.HitsCount == 0 {
			word.HitsCount = word.Progress
		}
	}

	fmt.Println(ResultMessage("Dictionary data is valid!"))
	fmt.Println()

	return StateAskNewWords{}, nil
}

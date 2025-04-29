package state

import (
	"fmt"
	"words/internal/model"
)

type StateProcessLessonResults struct {
	Words []Word
}

func (s StateProcessLessonResults) Execute(sm *StateMachine) (State, error) {
	fmt.Println(ProcessMessage("Processing results..."))

	for _, v := range s.Words {
		i, err := findDictionaryRecord(sm.ActiveWords, v.Word)
		if err != nil {
			return nil, err
		}

		pde := &sm.ActiveWords[i]
		pde.HitsCount += 1
		if v.Result == ResultCorrect {
			pde.Progress += 1
		} else if v.Result == ResultIncorrect && pde.Progress > 0 {
			pde.Progress -= 1
		}
	}

	fmt.Println(ResultMessage("Dictionary has been updated with the current progress..."))
	fmt.Println()

	return StateArchiveLearned{}, nil
}

func findDictionaryRecord(d []model.WordEntry, w string) (int, error) {
	for i := range d {
		if d[i].Word == w {
			return i, nil
		}
	}

	return -1, fmt.Errorf("the word '%v' is not found in the active dictionary", w)
}

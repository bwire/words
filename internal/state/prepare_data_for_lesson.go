package state

import (
	"fmt"
	"math/rand"
	"words/internal/config"
	"words/internal/model"
)

const (
	ResultNotAnswered Result = 0  // not checked yet
	ResultCorrect     Result = 1  // right answer
	ResultIncorrect   Result = -1 // wrong answer
)

type StatePrepareDataForLesson struct{}

type Result int

type Word struct {
	Word        string
	Translation string
	Progress    int
	Result      Result
}

func (s StatePrepareDataForLesson) Execute(sm *StateMachine) (State, error) {
	fmt.Println(ProcessMessage("Preparing data for lesson..."))

	selectedSet, err := PrepareDataForLesson(sm.ActiveWords)
	if err != nil {
		return nil, err
	}

	fmt.Println(ResultMessage("Words list is ready for lesson..."))
	fmt.Println()

	return StateLesson{Words: selectedSet}, nil
}

func PrepareDataForLesson(wordList []model.WordEntry) ([]Word, error) {
	// shuffle indexes
	indexes := make([]int, len(wordList))
	for i := range indexes {
		indexes[i] = i
	}

	rand.Shuffle(len(indexes), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})

	selectedSet := make([]Word, 0, config.WordsInSet)
	for i := 0; i < config.WordsInSet; i++ {
		idx := indexes[i]
		selectedSet = append(selectedSet, mapWordEntryToActiveWord(wordList[idx]))
	}

	return selectedSet, nil
}

func mapWordEntryToActiveWord(we model.WordEntry) Word {
	return Word{
		Word:        we.Word,
		Translation: we.Meaning,
		Progress:    we.Progress,
		Result:      ResultNotAnswered,
	}
}

package state

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"words/internal/config"
	"words/internal/model"
)

const (
	msgInputError      = "input error: %w"
	msgCancelledAdding = "Cancelled adding new words. Starting a new lesson"
)

type StateAskNewWords struct{}

func (s StateAskNewWords) Execute(sm *StateMachine) (State, error) {
	reader := bufio.NewReader(os.Stdin)

	ok, err := askYesNo(reader, "Would you like to add a new phrase (y/n)?")
	if err != nil {
		return nil, err
	}
	if !ok {
		fmt.Println()
		return StatePersistResult{}, nil
	}

	enWord, err := askWord(reader, "Enter a new english word (phrase):")
	if err != nil {
		fmt.Println(msgCancelledAdding)
		return StatePersistResult{}, nil
	}

	ruWord, err := askWord(reader, "Enter the meaning of the new word (phrase):")
	if err != nil {
		fmt.Println(msgCancelledAdding)
		return StatePersistResult{}, nil
	}

	newEntry := model.WordEntry{
		Word:      enWord,
		Meaning:   ruWord,
		StartDate: time.Now().Format(config.DateFormat),
		HitsCount: 0,
		Progress:  0,
	}

	sm.ActiveWords = append(sm.ActiveWords, newEntry)
	return StatePersistResult{}, nil
}

func askYesNo(reader *bufio.Reader, question string) (bool, error) {
	fmt.Println(question)
	for {
		answer, err := reader.ReadString('\n')
		if err != nil {
			return false, fmt.Errorf(msgInputError, err)
		}

		ok, err := normalizeYesNo(answer)
		if err == nil {
			return ok, nil
		}

		fmt.Println("Bad input! Should be y/n (or yes/no). Try once more, please:")
	}
}

func askWord(reader *bufio.Reader, prompt string) (string, error) {
	for {
		fmt.Println(prompt)
		word, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf(msgInputError, err)
		}
		word = strings.TrimSpace(word)

		if word != "" {
			return word, nil
		}

		ok, err := askYesNo(reader, "Empty input detected. Do you want to cancel adding new words? (y/n)")
		if err != nil {
			return "", err
		}
		if ok {
			return "", errors.New("cancelled by user. Starting a new lesson...")
		}
	}
}

func normalizeYesNo(answer string) (bool, error) {
	answer = strings.ToLower(strings.TrimSpace(answer))
	switch answer {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		return false, fmt.Errorf("bad input")
	}
}

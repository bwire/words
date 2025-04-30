package state

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"words/internal/config"
	"words/internal/model"
)

type StateArchiveLearned struct{}

func (s StateArchiveLearned) Execute(sm *StateMachine) (State, error) {
	fmt.Println(ProcessMessage("Checking words ready to archive..."))

	var newActiveWords []model.WordEntry
	var archiveWords []model.ArchiveEntry

	for _, w := range sm.ActiveWords {
		if w.Progress < 10 {
			newActiveWords = append(newActiveWords, w)
		} else {
			archiveWords = append(archiveWords, mapWordEntryToArchiveEntry(w))
		}
	}

	err := appendToArchive(archiveWords)
	if err != nil {
		return nil, err
	}

	fmt.Println(ProcessMessage("Archiving processed finished..."))
	fmt.Println()
	return StatePersistResult{NewDictionary: newActiveWords}, nil
}

func appendToArchive(archiveWords []model.ArchiveEntry) error {
	f, err := os.OpenFile(config.Archive, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)

	if len(archiveWords) > 0 {
		for _, aw := range archiveWords {
			err := encoder.Encode(aw)
			if err != nil {
				return fmt.Errorf("failed to append word to archive: %w", err)
			}
			fmt.Printf(
				ResultMessage("Word (phrase) '%v' is learned and goes to the archive! It took %v attempts to finish.\n"),
				aw.Word,
				aw.HitsCount,
			)
		}
	} else {
		fmt.Println(ResultMessage("No words to be archived!"))
	}

	return nil
}

func mapWordEntryToArchiveEntry(we model.WordEntry) model.ArchiveEntry {
	return model.ArchiveEntry{
		Word:        we.Word,
		Meaning:     we.Meaning,
		StartDate:   we.StartDate,
		HitsCount:   we.HitsCount,
		ArchiveDate: time.Now().Format(config.DateFormat),
	}
}

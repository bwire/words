package state

import (
	"fmt"
	"words/internal/model"
)

type State interface {
	Execute(sm *StateMachine) (State, error)
}

type StateMachine struct {
	current     State
	ActiveWords []model.WordEntry
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		current: StateLoadWords{},
	}
}

func (sm *StateMachine) Run() error {
	for sm.current != nil {
		nextStep, err := sm.current.Execute(sm)
		if err != nil {
			return fmt.Errorf("execution failed: %w", err)
		}
		sm.current = nextStep
	}
	return nil
}

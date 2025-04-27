package state

import "fmt"

type StateAskNewWords struct{}

func (s StateAskNewWords) Execute(sm *StateMachine) (State, error) {
	// TODO - rework
	for idx, we := range sm.ActiveWords {
		fmt.Printf("%v: %v\n", idx, we)
	}
	return nil, nil
}

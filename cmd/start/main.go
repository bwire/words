package main

import (
	"fmt"
	"words/internal/state"
)

func main() {
	fmt.Println("Starting vocabulary trainer...")

	machine := state.NewStateMachine()
	if err := machine.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

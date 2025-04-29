package main

import (
	"fmt"
	"os"
	"words/internal/config"
	"words/internal/state"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func initLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func checkError(err error, userMsg string, techMsg string) {
	if err != nil {
		fmt.Println(config.ColorError + userMsg + config.ColorReset)
		log.Error().Err(err).Msg(techMsg)
		os.Exit(1)
	}
}

func main() {
	initLogger()
	log.Info().Msg("Logger initialized successfully...")

	log.Info().Msg("Starting vocabulary trainer...")

	machine := state.NewStateMachine()
	checkError(machine.Run(), "working flow error", "state machine run failed")
}

package state

import "words/internal/config"

func ProcessMessage(m string) string {
	return colorize(m, config.ColorProcessMessage)
}

func ResultMessage(m string) string {
	return colorize(m, config.ColorResultMessage)
}

func ErrorMessage(m string) string {
	return colorize(m, config.ColorError)
}

func PromptWordMessage(m string) string {
	return colorize(m, config.ColorPrompt)
}

func colorize(m string, color string) string {
	return color + m + config.ColorReset
}

package config

const (
	// DateFormat defines date representation in each operation
	DateFormat = "2006-01-02"

	// WordsInSet - the minimum amount of words required in the dictionary to start a session
	WordsInSet = 10

	// Dictionary file name
	Vocabulary = "words.json"

	// Archive file name
	Archive = "archive"

	// Color scheme
	ColorReset          = "\033[0m"
	ColorProcessMessage = "\033[34m"
	ColorError          = "\033[31m"
	ColorResultMessage  = "\033[32m"
	ColorPrompt         = "\033[35m"
)

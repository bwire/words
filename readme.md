# words

A lightweight personal vocabulary trainer for learning English words and phrases.
This is a simple Go console application that tracks progress, handles repetition, and supports custom learning workflows.

## Installation

```bash
git clone https://github.com/your-username/words.git
cd words
```

## Usage

```bash
go run main.go
```

### Basic flow

1. Load your `word_list_export.json`
2. Start a vocabulary session in the console
3. Answer prompts and track your progress
4. Words that reach a score of 10 are automatically handled as "mastered"

## WordEntry structure

Each vocabulary item in `word_list_export.json` should follow this format:

```json
{
  "word": "to tug",
  "translation": "дёргать, тянуть",
  "progress": 4
}
```

- `word`: the English word or phrase to learn
- `translation`: its meaning in your native language
- `progress`: number of successful repetitions (0–10)

## Notes

- The app works entirely offline and uses JSON files for persistence.
- Mastered words are exported separately — no manual cleanup required.
- You can edit the word list manually or extend the logic to support UI or other file formats.

## License

This project is licensed under the MIT License.

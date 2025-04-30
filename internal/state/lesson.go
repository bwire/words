package state

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type StateLesson struct {
	Words []Word
}

func (s StateLesson) Execute(sm *StateMachine) (State, error) {
	fmt.Println("Starting lesson...")

	reader := bufio.NewReader(os.Stdin)

	for i := range s.Words {
		v := &s.Words[i]

		fmt.Printf(PromptWordMessage("%v: %v\n"), i+1, v.Translation)
		answer, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf(msgInputError, err)
		}

		if strings.TrimSpace(answer) == v.Word {
			v.Result = ResultCorrect
			fmt.Printf(ResultMessage("Correct answer!. The progress is %v!\n\n"), v.Progress+1)
		} else {
			v.Result = ResultIncorrect
			var progress = v.Progress
			if progress > 0 {
				progress -= 1
			}
			fmt.Printf(ErrorMessage("This is not correct! The correct answer is '%v'! Progress goes down (%v)!\n\n"), v.Word, progress)
		}
	}

	fmt.Println(ResultMessage("The lesson is over!"))
	fmt.Println()

	return StateProcessLessonResults(s), nil
}

package utils

const LT_COMMAND_STOP = 0
const LT_COMMAND_DDOS = 1
const LT_COMMAND_KEEP = 2

type Command struct {
	Id      int
	Command int
	Params  []string
}

func IsCorrectCommand(cmd Command) (isCorrect bool) {

	return true
}

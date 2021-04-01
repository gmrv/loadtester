package utils

const LT_COMMAND_STOP = 0
const LT_COMMAND_DDOS = 1
const LT_COMMAND_KEEP = 2

type Command struct {
	Id      int
	Command int
	Params  []string
}

type LTCommandError struct {}

var commandsDict []string
var commandsMap = make(map[string]int)

func IsCorrectCommand(cmd Command) (isCorrect bool) {
	var settings = GetSettings()
	var commands = settings.Commands

	for _, c := range commands {
		commandsDict = append(commandsDict, c.Command)
		commandsMap[c.Command] = len(c.Params)
	}

	return false
}

func (e LTCommandError) Error() string{return "123"}

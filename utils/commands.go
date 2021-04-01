package utils

const LT_COMMAND_STOP = 0
const LT_COMMAND_DDOS = 1
const LT_COMMAND_KEEP = 2

var commandsDict []string
var commandsMap = make(map[string]int)

type ParamType struct {
	Param       string `json:"param"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type CommandType struct {
	Id          int         `json:"id"`
	Command     string      `json:"command"`
	Description string      `json:"description"`
	Params      []ParamType `json:"params"`
}

type Command struct {
	Id      int
	Command int
	Params  []string
}

type LTCommandError struct{}

func IsCorrectCommand(cmd Command) (isCorrect bool) {
	var commands = Helper.Settings.Commands

	for _, c := range commands {
		commandsDict = append(commandsDict, c.Command)
		commandsMap[c.Command] = len(c.Params)
	}

	return false
}

func (e LTCommandError) Error() string { return "123" }

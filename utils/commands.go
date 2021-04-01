package utils

import "fmt"

const LT_COMMAND_STOP = "stop"
const LT_COMMAND_DDOS = "ddos"
const LT_COMMAND_KEEP = "keep"

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
	Params []string
}

type LTCommandError struct{
	command CommandType
}

func (e LTCommandError) Error() string {
	return fmt.Sprintf("Uncknown command \"%s\"", e.command.Name)
}

func IsCorrectCommand(cmd CommandType) (isCorrect bool, err error) {
	var commands = Helper.Settings.Commands

	for _, c := range commands {
		commandsDict = append(commandsDict, c.Name)
		commandsMap[c.Name] = len(c.Params)
	}

	if _, ok := commandsMap[cmd.Name]; ok {
		return true, nil
	}

	return false, LTCommandError{cmd}
}

func PutCommand(){
	Helper.CommandsStack = append(Helper.CommandsStack, CommandType{100, "ddos", nil})
	Helper.CommandsStack = append(Helper.CommandsStack, CommandType{101, "ddos", nil})
	Helper.CommandsStack = append(Helper.CommandsStack, CommandType{102, "ddos", nil})
}

func GetCommand(){
	Helper.CommandsStack = append(Helper.CommandsStack[:1], Helper.CommandsStack[2:]...)
}



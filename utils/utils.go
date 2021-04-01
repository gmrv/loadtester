package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Helper = GetHelper()

type Settings struct {
	Url              string        `json:"url"`
	RequPerRoutine   int           `json:"requests"`
	NumberOfRoutines int           `json:"routines"`
	Seconds          int           `json:"seconds"`
	Commands         []CommandType `json:"commands"`
}

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

func GetSettings() Settings {
	jsonFile, err := os.Open("settings.json")
	Check(err)

	WriteLog("Successfully Opened settings.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	Check(err)

	var settings Settings
	json.Unmarshal(byteValue, &settings)

	WriteLog(settings)
	return settings
}

func Check(e error) {
	if e != nil {
		WriteLog(e)
		panic(e)
	}
}

func WriteLog(v ...interface{}) {
	Helper.Logger.Print(v...)
	fmt.Println(v...)
}

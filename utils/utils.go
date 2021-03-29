package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Helper = GetHelper()

const SIG_STOP = 1

type Settings struct {
	Url              string `json:"url"`
	RequPerRoutine   int    `json:"requests"`
	NumberOfRoutines int    `json:"routines"`
	Seconds          int    `json:"seconds"`
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

func WriteLog(v ... interface{}){
	Helper.Logger.Print(v ...)
	fmt.Println(v ...)
}

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Settings struct {
	Url   				string 	`json:"url"`
	RequPerRoutine   	int 	`json:"requests"`
	NumberOfRoutines  	int    	`json:"routines"`
	Seconds				int		`json:"seconds"`
}

func GetSettings() (Settings) {
	jsonFile, err := os.Open("settings.json")
	Check(err)

	fmt.Println("Successfully Opened settings.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	Check(err)

	var settings Settings
	json.Unmarshal(byteValue, &settings)

	fmt.Println(settings)
	return settings
}

func LogResult(status int){
	fmt.Printf("%d;\n" ,status)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

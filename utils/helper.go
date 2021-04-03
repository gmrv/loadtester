// Тут будут объекты инстанс которых должен быть в единственном числе
// Логер
// Настройки
// Стек команд 

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var once sync.Once
var singleInstance *helper

type helper struct {
	logger log.Logger
	settings SettingsType
	CommandStack []CommandType
}

type SettingsType struct {
	Url              string                  `json:"url"`
	RequPerRoutine   int                     `json:"requests"`
	NumberOfRoutines int                     `json:"routines"`
	Seconds          int                     `json:"seconds"`
	Commands         []CommandDescriptorType `json:"commands"`
}

func GetHelper() *helper {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &helper{
					logger:        getLogger(),
					settings:      getSettings(),
				}
			})
	}
	return singleInstance
}

func getLogger() (log log.Logger) {
	file, err := os.OpenFile("loadtester.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.SetOutput(file)
	return log
}

func getSettings() SettingsType {
	jsonFile, err := os.Open("settings.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var settings SettingsType
	json.Unmarshal(byteValue, &settings)

	return settings
}


package utils

import (
	"fmt"
	"log"
)

var logger = GetLogger()

func Check(e error) {
	if e != nil {
		WriteLog(e)
		panic(e)
	}
}

func WriteLog(v ...interface{}) {
	logger.Print(v...)
	fmt.Println(v...)
}

func GetSettings() SettingsType {
	return GetHelper().settings
}

func GetLogger() log.Logger {
	return GetHelper().logger
}


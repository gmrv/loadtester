package utils

import (
	"fmt"
)

var Helper = GetHelper()

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

// Тут будут объекты инстанс которых должен быть в единственном числе
// Логер

package utils

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var once sync.Once

type helper struct {
	Logger log.Logger
}

var singleInstance *helper

func GetHelper() *helper {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &helper{Logger: getLogger()}
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

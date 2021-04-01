package main

import (
	"loadtester/utils"
	"net/http"
	"strconv"
	"time"
)

var keepGoing = true
var counter = 0
var settings = utils.GetSettings()

func performRequest(url string) {
	resp, err := http.Get(url)
	if err == nil {
		utils.WriteLog(resp.StatusCode, " ", url)
		resp.Body.Close()
	} else {
		utils.WriteLog(err)
	}
}

func performMultiRequest(url string, count int) {
	if count == 0 {
		for {
			if !keepGoing {
				break
			}
			performRequest(url)
		}
	} else {
		for i := 0; i < count; i++ {
			if !keepGoing {
				break
			}
			performRequest(url)
		}
	}
}

func getCommand() (command utils.CommandType, err error) {
	var cmd = utils.CommandType{}
	counter += 1

	// Получаем команду

	// Проверяем команду

	// Записываем ее в стек комманд хелпера (FIFO)

	// Из шины потом прочитаем в эту функцию

	// Считаем количество секунд и посылаем команду на завершение если время истекло
	if counter >= settings.Seconds {
		return utils.CommandType{2, utils.LT_COMMAND_STOP, nil}, err
	}

	cmd = utils.CommandType{1, utils.LT_COMMAND_DDOS, nil}

	ok, err := utils.IsCorrectCommand(cmd)
	if !ok {
		panic(err)
	}

	return cmd, err
}

func main() {
	var last_command_id = -1

	utils.WriteLog("Start.")

	var h = utils.Helper
	utils.WriteLog(h)

	utils.PutCommand()
	utils.GetCommand()

	for {

		command, err := getCommand()
		utils.Check(err)

		// Не выполнять уже выполненную команду
		if last_command_id != command.Id {

			last_command_id = command.Id

			switch command.Name {
			case utils.LT_COMMAND_DDOS:
				for i := 0; i < settings.NumberOfRoutines; i++ {
					go performMultiRequest(settings.Url+"?a="+strconv.Itoa(i), settings.RequPerRoutine)
				}

			case utils.LT_COMMAND_STOP:
				keepGoing = false

			default:
				keepGoing = true
			}

		}

		if !keepGoing {
			break
		}

		time.Sleep(1000 * time.Millisecond)

	}

	utils.WriteLog("Finished.")
}

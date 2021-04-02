package main

import (
	"fmt"
	"loadtester/utils"
	"net/http"
	"strconv"
	"time"
)

var keepGoing = true
var counter = 0
var settings = utils.GetSettings()
var requests = 0

func performRequest(url string) (status int) {
	status = 0
	resp, err := http.Get(url)
	if err == nil {
		status = resp.StatusCode
		utils.WriteLog(status, " ", url)
		resp.Body.Close()
	} else {
		status = -1
		utils.WriteLog(err)
	}
	return status
}

func performMultiRequest(url string, count int, c chan int) {
	if count == 0 {
		for {
			if !keepGoing {
				break
			}
			c <- performRequest(url)
		}
	} else {
		for i := 0; i < count; i++ {
			if !keepGoing {
				break
			}
			c <- performRequest(url)
		}
	}
}

func listenChanel(c chan int)  {
	for {
		<-c
		requests += 1
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

	c := make(chan int, 100)

	go listenChanel(c)

	for {

		command, err := getCommand()
		utils.Check(err)

		// Не выполнять уже выполненную команду
		if last_command_id != command.Id {

			last_command_id = command.Id

			switch command.Name {
			case utils.LT_COMMAND_DDOS:
				for i := 0; i < settings.NumberOfRoutines; i++ {
					go performMultiRequest(settings.Url+"?a="+strconv.Itoa(i), settings.RequPerRoutine, c)
				}

			case utils.LT_COMMAND_STOP:
				keepGoing = false

			default:
				keepGoing = true
			}

		}

		if !keepGoing {
			utils.WriteLog("wait...")
			time.Sleep(5000 * time.Millisecond)
			break
		}

		time.Sleep(1000 * time.Millisecond)

	}

	defer fmt.Println(">>>", requests)
	utils.WriteLog("Finished.")
}

package main

import (
	"fmt"
	"loadtester/utils"
	"time"
)

var counter = 0
var settings = utils.GetSettings()
var requests = 0

func listenChanel() {
	for {
		<-utils.GetHelper().C
		requests += 1
	}
}

func getCommand() (command utils.LTCommand, err error) {
	counter += 1

	// Получаем команду

	// Проверяем команду

	// Записываем ее в стек комманд хелпера (FIFO)

	// Из шины потом прочитаем в эту функцию

	// Считаем количество секунд и посылаем команду на завершение если время истекло
	if counter >= settings.Seconds {
		cmd := &utils.LTCommandStop{
			Id: 2,
		}
		return cmd, err
	}

	cmd := &utils.LTCommandDDOS{
		Id:       1,
		Url:      settings.Url,
		Routines: settings.NumberOfRoutines,
		Requests: settings.RequPerRoutine,
		Seconds:  settings.Seconds,
	}

	//ok, err := utils.IsCorrectCommand(cmd)
	//if !ok {
	//	panic(err)
	//}

	return cmd, err
}

func main() {
	var last_command_id = -1

	utils.WriteLog("Start.")

	go listenChanel()

	for {

		command, err := getCommand()
		utils.Check(err)

		// Не выполнять уже выполненную команду
		if last_command_id != command.GetId() {

			last_command_id = command.GetId()

			utils.WriteLog("Command ", command)

			command.Execute()

		}

		if utils.GetHelper().FL_QUIT {
			break
		}

		time.Sleep(1000 * time.Millisecond)

	}

	defer fmt.Println(">>>", requests)
	utils.WriteLog("Finished.")
}

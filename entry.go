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

func getCommand() (command int, params []string, err error) {
	//if command == utils.LD_COMMAND_STOP {
	//	keepGoing = false
	//}
	counter += 1
	if counter >= settings.Seconds{
		return utils.LD_COMMAND_STOP, nil, nil
	}
	if counter != 1 {
		return utils.LD_COMMAND_KEEP, nil, nil
	}

	return utils.LD_COMMAND_DDOS, nil, nil
}

func main() {
	utils.WriteLog("Start.")

	for  {

		command, params, err := getCommand()
		utils.Check(err)

		switch command {

		case utils.LD_COMMAND_DDOS:
			for i := 0; i < settings.NumberOfRoutines; i++ {
				go performMultiRequest(settings.Url+"?a="+strconv.Itoa(i), settings.RequPerRoutine)
			}
			fmt.Println(params)

		case utils.LD_COMMAND_STOP:
			keepGoing = false

		case utils.LD_COMMAND_KEEP:
			keepGoing = true

		default:
			keepGoing = true
		}

		if !keepGoing{
			break
		}

		time.Sleep(1000 * time.Millisecond)

	}

	utils.WriteLog("Finished.")
}

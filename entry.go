package main

import (
	"loadtester/utils"
	"net/http"
	"strconv"
	"time"
)

var keepGoing = true

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

func getCommand() {
	var command = utils.LD_COMMAND_STOP

	if command == utils.LD_COMMAND_STOP {
		keepGoing = false
	}
}

func main() {
	utils.WriteLog("Start.")

	var settings = utils.GetSettings()

	for i := 0; i < settings.NumberOfRoutines; i++ {
		go performMultiRequest(settings.Url+"?a="+strconv.Itoa(i), settings.RequPerRoutine)
	}

	for i := 0; i < settings.Seconds; i++ {
		time.Sleep(1000 * time.Millisecond)
		//getCommand()
	}

	utils.WriteLog("Finished.")
}

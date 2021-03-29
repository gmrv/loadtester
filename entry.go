package main

import (
	"fmt"
	"loadtester/utils"
	"net/http"
	"strconv"
	"time"
)

var keepGoing = true

func performRequest(url string) {
	//var client http.Client
	resp, err := http.Get(url)
	if err == nil {
		fmt.Println(resp.StatusCode)
		resp.Body.Close()
	} else {
		fmt.Println(err)
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

	var settings = utils.GetSettings()

	for i := 0; i < settings.NumberOfRoutines; i++ {
		go performMultiRequest(settings.Url+"?a="+strconv.Itoa(i), settings.RequPerRoutine)
	}

	for i := 0; i < settings.Seconds; i++ {
		fmt.Println(">>>>>>>>>>>", i)
		time.Sleep(1000 * time.Millisecond)

		getCommand()

	}

	fmt.Println("Finished.")
}

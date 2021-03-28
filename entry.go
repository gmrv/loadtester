package main

import (
	"fmt"
	"loadtester/utils"
	"net/http"
	"strconv"
	"time"
)

func performRequest(url string){
	//var client http.Client
	resp, err := http.Get(url)
	if err == nil{
		fmt.Println(resp.StatusCode)
		resp.Body.Close()
	}else{
		fmt.Println(err)
	}
}

func performMultiRequest(url string, count int){
	if count == 0{
		for{
			performRequest(url)
		}
	}else {
		for i := 0; i < count; i++{
			performRequest(url)
		}
	}
}

func main() {

	var settings = utils.GetSettings()

	for i := 0; i < settings.NumberOfRoutines; i++ {
		go performMultiRequest(settings.Url + "?a=" + strconv.Itoa(i), settings.RequPerRoutine)
	}

	for i := 0; i < settings.Seconds; i++ {
		time.Sleep(1000 * time.Millisecond)
		//
	}

	fmt.Println("Finished.")
}

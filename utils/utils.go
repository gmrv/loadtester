package utils

import (
	"fmt"
)

func LogResult(status int){
	fmt.Printf("%d;\n" ,status)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

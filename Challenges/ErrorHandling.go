package Challenges

import (
	"fmt"
	"strconv"
)

func TestErrorHandling() {
	handleStringConversionError()
}

func handleStringConversionError() {
	var string1 string
	var int1 int
	var err error

	string1 = "12ABC"

	int1, err = strconv.Atoi(string1)

	if err != nil {
		fmt.Println("Error is ", err)
	} else {
		fmt.Printf("Converted number %d\n", int1)
	}
}

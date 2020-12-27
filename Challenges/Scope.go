package Challenges

import "fmt"

var a1 string = "G"

func TestGlobal() {
	fmt.Println("a1 ", a1)
	a1 := "Local a1"
	fmt.Println("a1", a1)

	useGlobal()
}

func useGlobal() {
	fmt.Println("Global a1", a1)
}

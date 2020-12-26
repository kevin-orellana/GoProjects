package main

import (
	"HelloWorld/Challenges"
	"fmt"
)

func main(){
	fmt.Println("Hello World!!!")
	//你好(); // Go identifiers need only be Unicode-standard

	// Challenges
	Challenges.Runner()
	//Challenges.TestGlobal()
	//Challenges.TestControlStructures()
	//Challenges.TestErrorHandling()
	Challenges.TestSimpleFunctions()
}

func 你好(){
	fmt.Println("calling another func")
	var number1 float32 = 5.2
	fmt.Println(number1)
	fmt.Println(int(number1)) // float32 explicitly casted to int

	fmt.Println("first hour ", hour(first), "second hour", hour(second))

	var number int          // Declaring  an integer variable
	fmt.Println(number)     // Printing its value
	var decision bool       // Declaring a boolean variable
	fmt.Println(decision)   // Printing its value

	number2 := 4 // shortform intialization
	fmt.Println("short form initialization", number2)
	boolean2 := true
	fmt.Printf( "printf %d, %t", number2, boolean2)
}

type hour int
const (
	first = iota
	second
	third
)
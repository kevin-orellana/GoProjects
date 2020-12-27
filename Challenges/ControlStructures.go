package Challenges

import "fmt"

func TestControlStructures() {
	//ifElseStructures()
	//switchCaseExample()
	//forLoopExample()
	//breakContinueExample()
	//labelsExample()
	//goToExample()
}

func ifElseStructures() {
	var condition1 bool = true
	var condition2 bool = true

	if condition1 {
		println("first if block")
	} else if condition2 {

	} else {

	}
	/* Notes
	- observe no line breaks between conditional statements and braces
	- observe no parentheses outside conditional statements
	*/
	max := 5
	if val := 10; val > max {
		fmt.Println("Test initializing in the if structure ")
	}
}

func switchCaseExample() {
	var num1 int = 100
	switch num1 {
	case 1, 2, 3:
		{
			fmt.Println("1, 2 or 3")
		}
	case 20:
		fallthrough
	case 100:
		{
			fmt.Println("value is ", num1)
		}
	default:
		{
			fmt.Println("default case")
		}
	}

	switch num := 10; {
	case num1 < 0:
		{
			println("number is negative")
		}
	case num > 0 && num < 10:
		{
			println("number is between 0 and 10")
		}
	default:
		{
			println("numb is greater than 10")
		}
	}
}

func forLoopExample() {
	for i := 0; i < 2; i++ {
		fmt.Println("i", i)
	}

	var str1 string = "Hello"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%d-%c ", i, str1[i]) // will iterate through each byte in str1
	}
	fmt.Println()

	var str2 string = "你好大家"
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%d-%c  ", i, str2[i]) // will iterate through each byte in str1 which in this case may be 4 times 4
	}
	fmt.Println()

	// for range example used to solve iterating through a Unicode str iteration
	for pos, char := range str2 {
		fmt.Printf("Char at %d is %c", pos, char)
	}
	fmt.Println()
}

func breakContinueExample() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break // breaks the inner most nested loop that iterates on j
			}

			if j == 3 {
				continue // will continue on with the j loop (and print 4)
			}
			fmt.Printf("%d ", j)
		}
	}
}

func labelsExample() {
LABEL1: // add a label for location
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1 // jump to the label -- essentially jump to next iteration of outer loop marked by LABEL1
			}
			fmt.Printf("i is %d and j is %d\n", i, j)
		}
	}
}

func goToExample() {
	i := 0
IDENTIFIER1GOTO: // mark with colon
	i++
	if i == 5 {
		fmt.Println("Exiting goToExaple")
		return
	}
	fmt.Printf("jumping when i is %d\n ", i)
	goto IDENTIFIER1GOTO // goto IDENTIFIER1GOTO
}

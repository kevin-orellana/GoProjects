package Challenges

import "fmt"

// aliasing
type Celcius float32
type Fahrenheit float32

func Runner() {
	fmt.Println("Running Temperature Conversion Challenge")

	// Temp conversion
	var tCelcius1 Celcius = 100
	fmt.Println("Celcius ", tCelcius1, "Farenheit", toFarenheit(tCelcius1))

	// Season name
	var month1 int = 10
	fmt.Println("Month", month1, "has season", seasonOfTheMonth(month1))

	//computeFactorialOfANumber()
	testFilteringEvenAndOdds()
}

func toFarenheit(tCelcius Celcius) Fahrenheit {
	var tFarenheit Fahrenheit = Fahrenheit(tCelcius*(9.0/5.0) + 32)
	return tFarenheit
}

func seasonOfTheMonth(monthNumber int) string {
	switch monthNumber {
	case 1, 2, 12:
		{
			return "Winter"
		}
	case 3, 4, 5:
		{
			return "Spring"
		}
	case 6, 7, 8:
		{
			return "Summer"
		}
	case 9, 10, 11:
		{
			return "Autumn"
		}
	default:
		{
			return "Unknown"
		}
	}
}

func computeFactorialOfANumber() {
	// write a function that returns the factorial (!) of n
	// the factorial of n is n * (n-1) * ... * (n-n)! where (0)! is 1

	var n uint64 = 10 // maximum is 21 because after this Overflow will occur
	fmt.Printf("factorial of %d is %d\n", n, factorial(n))
}

func factorial(input uint64) (output uint64) {
	if input == 0 {
		return 1
	}
	output = input * factorial(input-1)
	return // automatically returns 'output'
}

// filter even and odds using callback function
type filter func(int) bool // aliasing type
func testFilteringEvenAndOdds() {
	var slice []int = []int{1, 2, 3, 4, 5, 7}
	evens, odds := performFilter(slice, isEven)
	fmt.Println("evens", evens, "odds", odds)
}

func performFilter(slice []int, function1 filter) (evens, odds []int) {
	for v, _ := range slice {
		if function1(v) {
			evens = append(evens, v)
		} else {
			odds = append(odds, v)
		}
	}
	return
}

func isEven(int1 int) (isEven bool) { // adheres to the filter type
	isEven = int1%2 == 0
	return
}

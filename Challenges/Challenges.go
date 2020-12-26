package Challenges
import "fmt"


// aliasing
type Celcius float32
type Fahrenheit float32

func Runner(){
	fmt.Println("Running Temperature Conversion Challenge")

	// Temp conversion
	var tCelcius1 Celcius = 100
	fmt.Println("Celcius ", tCelcius1, "Farenheit", toFarenheit(tCelcius1))

	// Season name
	var month1 int = 10;
	fmt.Println("Month", month1, "has season", seasonOfTheMonth(month1))

}

func toFarenheit(tCelcius Celcius) Fahrenheit {
	var tFarenheit Fahrenheit = Fahrenheit(tCelcius * (9.0/5.0) + 32)
	return tFarenheit
}

func seasonOfTheMonth(monthNumber int) string{
	switch monthNumber {
	case 1, 2, 12:{
		return "Winter"
	}
	case 3, 4, 5:{
		return "Spring"
	}
	case 6, 7, 8:{
		return "Summer"
	}
	case 9, 10, 11:{
		return "Autumn"
	}
	default:{
		return "Unknown"
	}
	}
}

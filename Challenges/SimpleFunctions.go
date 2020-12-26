package Challenges

import "fmt"

func TestSimpleFunctions(){

	a1, b1 := 3, 4
	sum1, product1, difference1 := multipleValues(a1, b1)
	fmt.Printf("returned values: %d %d %d\n", sum1, product1, difference1)

	sum1 = variableNumberArguments(5, -2, 0, 9)
	fmt.Printf("sum is %d", sum1)

	deferExample()
}

func multipleValues(a int, b int) (sum int, product int, difference int) {
	return a + b, a * b, a - b
}

func variableNumberArguments(ints ...int) (sum int){
	for _, v := range ints { // _ is discarded
		sum += v
	}
	return // note how we just 'return' and the declared sum is returned
}

func deferExample() {
	fmt.Println("Doing something and returning")

	for i := 0; i < 2; i++ {
		defer fmt.Printf("Deferred int %d\n", i)
	}
	defer somethingToDefere()
	return
}

func somethingToDefere() {
	fmt.Println("Something to defer")
}
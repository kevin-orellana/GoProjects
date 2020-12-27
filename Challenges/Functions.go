package Challenges

import "fmt"

func TestSimpleFunctions() {

	a1, b1 := 3, 4
	sum1, product1, difference1 := multipleValues(a1, b1)
	fmt.Printf("returned values: %d %d %d\n", sum1, product1, difference1)

	sum1 = variableNumberArguments(5, -2, 0, 9)
	fmt.Printf("sum is %d", sum1)

	//deferExample()
	//builtInFunctions()
	//callbackExample()
	closuresExample()
}

// multiple values returned by a function and variable arguments
func multipleValues(a int, b int) (sum int, product int, difference int) {
	return a + b, a * b, a - b
}

func variableNumberArguments(ints ...int) (sum int) {
	for _, v := range ints { // _ is discarded
		sum += v
	}
	return // note how we just 'return' and the declared sum is returned
}

// using the defer keyword
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

func builtInFunctions() {
	var str1 string = "A String"
	fmt.Printf("String %s has len %d", str1, len(str1))

	v := new(int) // use new keyword to define a value tyupe
	fmt.Printf("%d has been initialized with new keyword", v)
}

// using callbacks and function types
func callbackExample() {
	add(1, 3, callbackAdd) // note: we pass the name of a function!
}

func callbackAdd(a, b int) { // used as a callback function
	fmt.Printf("CALLBACK: A is %d B is %d - sum is %d\n", a, b, a+b)
}

func add(a int, b int, f func(int, int)) { // note the function parameter
	callbackAdd(a, b)
}

// using closures (lambda, anonymous or literal functions)
func closuresExample() {
	// function literal example - fplus is assigned a function
	fplus := func(x, y int) int { return x + y }
	fmt.Println("Result of using fplus", fplus(3, 5))

	// you can reassigna  function literal
	for i := 0; i < 3; i++ {
		g := func(i int) int { fmt.Println("reassign fun literal value", i); return i }
		g(i)
		fmt.Printf("g is of type %T and has value %v\n", g, g)
	}
	// invoking a function literal directly
	func(x, y int) int {
		fmt.Println("literal function result", x+y)
		return x + y
	}(3, 4) // here we invoke the function literal directly

	// similarly
	func() {
		var sum float32 = 0.0
		for i := 1; i < 1e6; i++ {
			sum += float32(i)
		}
		fmt.Println("Result of function literal ", sum)
	}()

	incBy3 := generateFunction(3)
	fmt.Println("Result of generated func ", incBy3(2))
}

// example of returning a function using closures (function literal generation_
func generateFunction(n int) func(x int) (result int) {
	return func(x int) int {
		return n + x
	}
}

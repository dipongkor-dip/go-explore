package callbackFunctions

import "fmt"

// callback function
func process(function_name func()) {
	function_name()
}

func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

// higher order function
func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

func CallbackFunc() {
	// var fun = func() {
	// 	fmt.Println("Hello world!!")
	// }
	// process(fun)

	// var add = func(a int, b int) int {
	// 	return a + b
	// }
	// var multiply = func(a int, b int) int {
	// 	return a * b
	// }

	// fmt.Println("sum:", calculate(5, 3, add), "multiply:", calculate(5, 3, multiply))

	// anonymous callback function
	// var res = calculate(5, 3, func(x, y int) int { return x - y })
	// fmt.Println("subtract:", res)

	// higher order function
	var double = multiplyBy(2) //func(x int) int { return factor * x }
	fmt.Println("double val:", double(5), ",", double(100))

	var triple = multiplyBy(3) //func(x int) int { return factor * x }
	fmt.Println("triple val:", triple(5), ",", triple(100))

}

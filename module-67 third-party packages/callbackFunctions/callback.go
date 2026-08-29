package callbackFunctions

import "fmt"

// callback function
func process(function_name func()) {
	function_name()
}

func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

func CallbackFunc() {
	// var fun = func() {
	// 	fmt.Println("Hello world!!")
	// }
	// process(fun)

	var add = func(a int, b int) int {
		return a + b
	}
	var multiply = func(a int, b int) int {
		return a * b
	}

	fmt.Println("sum:", calculate(5, 3, add), "multiply:", calculate(5, 3, multiply))
}

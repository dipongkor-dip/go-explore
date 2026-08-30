package main

import "fmt"

// higher order function
func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

func makeCounter() func() int {
	var count int = 0

	var inner = func() int {
		count++
		return count
	}

	return inner
}

func clouser_func() {
	// var double = multiplyBy(2)
	// fmt.Println("double val", double(5))

	var next func() int = makeCounter()

	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3
}

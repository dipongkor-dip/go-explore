package main

import "fmt"

func array_func() {
	var numbers [6]int // type integer and size 6

	numbers[1] = 10

	numbers[0] = 20

	// fmt.Println(numbers) // [20 10 0 0 0 0]

	// fmt.Println("Length of numbers array: ", len(numbers))
	// alternative
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	var num = [6]int{10, 20, 30, 40}

	fmt.Println(num) // [10 20 30 40 0 0]
}

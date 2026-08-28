package main

import "fmt"

func slice_func() {
	var orders = [6]int{10, 20, 30, 40, 50, 60}

	var slice []int = orders[1:4]
	// fmt.Println(slice) // its look like a window

	// slice := orders[:]
	// fmt.Println(slice)

	// slice[0] = 100

	// fmt.Println(slice)
	// fmt.Println(orders)

	slice = append(slice, 101)

	// slice = append(slice, 103)
	// slice = append(slice, 104)

	fmt.Println(slice)

	fmt.Println("The length of the slice is: ", len(slice))
	fmt.Println("The capacity of the slice is: ", cap(slice))

}

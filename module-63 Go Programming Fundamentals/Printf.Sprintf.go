package main

import "fmt"

func print_sprint() {
	var name = "next level"

	age := 23

	ratting := 3.1416

	// fmt.Printf("Name : %s, Age: %d, Floating Value: %f \n", name, age, ratting)

	var str string = fmt.Sprintf("Name : %s, Age: %d, Floating Value: %f \n", name, age, ratting)

	fmt.Println(str)
}

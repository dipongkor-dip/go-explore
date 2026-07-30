package main

import "fmt"

// var name string = "dip"
// name := "dip" // wrong way outside of main func

func variables_in_go() {

	fmt.Println("Exploring Variables!")

	// variable declaration
	// var name string // default value "" empty string
	// var name string = "dip"
	// name := "dip" // short way
	// var name = "dip"

	// grouped variable declaration
	// var (
	// 	name string = "dip"
	// 	age  int    = 23
	// )

	// multiple variable declaration
	// var name, address string = "dip", "dhaka"
	// name = "dipongkor"
	// address = "dhaka"

	const name string = "dip"

	// var flag bool // default value false
	// var n int // default value 0

	fmt.Println(name)
}

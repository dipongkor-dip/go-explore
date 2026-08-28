package main

import "fmt"

type Employee struct {
	name       string
	age        int
	isLoggedIn bool
}

func receiver_struct_func() {
	var emp = Employee{
		name:       "John",
		age:        25,
		isLoggedIn: false,
	}

	emp.greet()

	// empPointer := &emp
	// empPointer.login()
	// go efficient alternative
	emp.login()

	fmt.Printf("%+v\n", emp)
}

func (e *Employee) greet() {
	fmt.Println("Hello", e.name)
}

func (e *Employee) login() {
	// (*e).isLoggedIn = true
	// go efficient alternative
	e.isLoggedIn = true
}

package main

import "fmt"

func doSomething() {
	defer func() {
		fmt.Println("deferred function run")
		var r = recover()
		if r != nil {
			fmt.Println("Recover panic from -", r)
		}
	}()
	panic("Something really bad happened")
}

func panic_func() {
	doSomething()

	fmt.Println("Main completed normally")
}

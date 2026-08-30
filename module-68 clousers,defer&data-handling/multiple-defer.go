package main

import "fmt"

func connectDb() {
	defer fmt.Println("Closing database connection ...")

	fmt.Println("Connecting to database ...")
}

func multiple_defer_func() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	connectDb()
}

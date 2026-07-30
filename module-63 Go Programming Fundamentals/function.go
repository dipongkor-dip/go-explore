package main

import "fmt"

// function
func user(name string) {
	fmt.Printf("Name : %s \n", name)
}

// function value return
func employee(name string) string {
	return fmt.Sprintf("Employee name: %s", name)
}

// multiple value return
func admin(name string, phone int) (string, int) {
	return name, phone
}

// multiple value return (with name)
func customer(name string, phone int) (cus_name string, cus_phone int) {
	cus_name = name // equal use because cus_name already declare
	cus_phone = phone

	// return // not need tell again
	return cus_name, cus_phone // (recommendation)
}

func functions_basics() {
	// user("sunny")

	// emp := employee("alo")
	// fmt.Println(emp)

	// name, phone := admin("Sunny", 2345)
	// fmt.Printf("Admin name: %s, phone: %d \n", name, phone)

	name, phone := customer("Sunny", 2345)
	fmt.Printf("Admin name: %s, phone: %d \n", name, phone)
}

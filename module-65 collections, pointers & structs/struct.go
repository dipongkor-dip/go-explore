package main

import "fmt"

type User struct {
	name           string
	email          string
	phone          int
	additionalInfo AdditionalInfo
}

// 1. Define a custom type for type-safety
type Status string

// 2. Define the allowed constants using that type
const (
	Single      Status = "single"
	Married     Status = "married"
	Complicated Status = "complicated"
)

type AdditionalInfo struct {
	age    int
	status Status
}

type Employee struct {
	name string
	age  int
	role string
}

func struct_func() {
	// var john = User{name: "John", email: "john1@gmail.com", phone: 123, additionalInfo: AdditionalInfo{age: 23, status: Single}}

	// fmt.Printf("%+v\n", john)

	// john.email = "john123@gmail.com"

	// fmt.Println(john.email)

	// var sunny User
	// sunny.name = "sunny"
	// sunny.email = "sunny@gmail.com"

	// fmt.Printf("%+v\n", sunny)

	emp := func(name string, age int, role string) Employee {
		if name == "" || age == 0 {
			fmt.Println("User name and age is required")
			return Employee{}
		}

		return Employee{
			name: name,
			age:  age,
			role: role,
		}
	}

	// 2. Call the function
	fmt.Println(emp("jamal", 25, "student"))

	// fmt.Print(emp("jamal", 25, "student"))
}

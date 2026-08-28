package main

import "fmt"

type user struct {
	name  string
	email string
	phone int
}

func struct_func() {
	var john = user{name: "John", email: "john1@gmail.com", phone: 123}

	fmt.Printf("%+v\n", john)

	john.email = "john123@gmail.com"

	fmt.Println(john.email)

	var sunny user
	sunny.name = "sunny"
	sunny.email = "sunny@gmail.com"

	fmt.Printf("%+v\n", sunny)
}

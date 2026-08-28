package main

type User struct {
	name       string
	age        int
	isLoggedIn bool
	greet      func()
}

func main() {
	// var user = User{
	// 	name:       "John",
	// 	age:        25,
	// 	isLoggedIn: false,
	// 	greet:      func() { fmt.Println("Hello world") },
	// }

	// user.greet = func() { fmt.Println("Hello", user.name) }

	// user.greet()

	// receiver_struct_func()
	// maps_func()
	// range_func()
	// interface_func()
	payment_service_func()
}

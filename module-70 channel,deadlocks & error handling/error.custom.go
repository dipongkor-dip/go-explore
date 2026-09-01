package main

import (
	"errors"
	"fmt"
	// "fmt"
)

type CustomError struct {
	message string
	code    int
}

func (cu *CustomError) Error() string {
	return cu.message
}

func login(pass string) error {
	if pass != "1234" {
		// return &CustomError{message: "Password do not match", code: 401}
		return errors.New("password do not match")
	}

	return nil
}

func custom_error_func() {
	var err = login("234")
	if err != nil {
		// fmt.Println("Error", err, "Code", err.code) // not use because type assertion CustomError

		if customError, ok := err.(*CustomError); ok {
			fmt.Println(customError.code)
		}

	} else {
		fmt.Println("Login successful")
	}

	fmt.Println("Main end")

	// var users = map[int]string{1: "sunny", 2: "abi", 3: "alo"}
	// var name, ok = users[4] // if found 4 key value, ok == true
	// if ok == true {
	// 	fmt.Println(name)
	// } else {
	// 	fmt.Println("Not found")
	// }
}

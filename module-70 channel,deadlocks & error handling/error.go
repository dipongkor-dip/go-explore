package main

import (
	"errors"
	"fmt"
)

func error_handler_func() {
	var res, error = divide(10, 0)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(res)

}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

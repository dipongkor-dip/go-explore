package main

import "fmt"

func add_number(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func greet(prefix string, users ...string) {
	for _, u := range users {
		fmt.Println(prefix, u)
	}
}

func variadic_func() {
	sum := add_number(5, 10, 20)
	fmt.Println(sum)

	mps := []string{"jamal", "alo"}

	// variadic arguments
	greet("Welcome", mps...)

}

// flexible amount of argument
// must be the last parameter
// internally slice

package main

import "fmt"

func loop_func() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Error reading input: ", err)
	}

	for i := 1; i <= n; i++ {

		if i%2 == 0 {
			continue
		}

		if i > 100 {
			break
		}

		fmt.Printf("%d ", i)

	}
	fmt.Println()
}

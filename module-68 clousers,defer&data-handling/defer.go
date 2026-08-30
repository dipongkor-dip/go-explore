package main

import "fmt"

func deferred(res int) {
	fmt.Println("deferred result f1:", res)
}

func example() int {
	var res = 10

	// third ~ defer
	defer func() {
		res += 200
		fmt.Println("deferred result f2:", res)
	}()

	// first
	fmt.Println("example function res:", res)

	// second
	res += 100

	return res // remember 110 then call defer func then return 110
}

// name return
func example_() (res int) {
	res = 10

	// third ~ defer
	defer func() {
		res += 200
		fmt.Println("deferred result f2:", res)
	}()

	// first
	fmt.Println("main function res:", res)

	// second
	res += 100

	return // no catch - evaluate res reference
}

func defer_func() {
	// var res int = example()
	// fmt.Println("return result:", res)

	// fmt.Println()

	// var res_name int = example_()
	// fmt.Println("return result:", res_name)
}

// example function res: 10
// deferred result f2: 310
// return result: 110

// example function res: 10
// deferred result f2: 310
// return result: 310

package main

import "fmt"

func change(x *int) {
	*x = 100
	fmt.Println("Inside function: ", *x)
}

func array_pointer(arr *[]int) {
	fmt.Printf("Array header address: %p\n", arr)
	(*arr)[1] = 101
	fmt.Println(*arr)
}

func array_copy(arr []int) {
	fmt.Printf("Array header address: %p\n", &arr)
	arr[1] = 101
	fmt.Println(arr)
}

func pointer_func() {
	var a int = 3
	// var p *int = &a
	// alternative
	p := &a

	*p = 5

	// fmt.Println("a:", a)  // 5
	// fmt.Println("b:", *p) // 5 ~ dereference

	// change(&a)

	// fmt.Println("a:", a)  // 5
	// fmt.Println("b:", *p) // 5 ~ dereference

	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("Array header address: %p\n", &array)

	array_pointer(&array)

	array_copy(array)

}

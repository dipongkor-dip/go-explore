package main

import "fmt"

// interface{} == any

// type assertion
// ok idiom

func Process(data any) {
	// fmt.Println(data)

	strData, ok := data.(string)
	if ok {
		fmt.Println("Data is string", strData)
	}

	intData, ok := data.(int)
	if ok {
		fmt.Println("Data is integer", intData+100)
	}

	boolData, ok := data.(bool)
	if !ok {
		fmt.Println("Data is not bool")
	} else {
		fmt.Println("Data is boolean", boolData)
	}
}

func type_handle_func() {
	var data interface{}

	data = "sunny"
	fmt.Println(data)

	data = 25
	fmt.Println(data)

	Process([]int{1, 2, 3, 4, 5})
	Process("Next level")

	Process(false)

}

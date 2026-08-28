package main

import "fmt"

func range_func() {
	// mp := map[string]string{"name": "sunny", "success": "ok"}
	// fmt.Printf("%v\n", mp)
	// for k, v := range mp {
	// 	fmt.Println(k, v)
	// }

	// for _, v := range mp {
	// 	fmt.Println(v)
	// }

	// colors := []string{"blue", "red"}
	// colors = append(colors, "yellow")
	// fmt.Println(colors)

	name := "NextLevel"
	// var byte_slice []byte // under the hood working as a byte slice

	var byteSlice = []byte(name)

	// for i, v := range name {
	// 	fmt.Println(i, v) // index , ascii value
	// }

	fmt.Println(byteSlice)
}

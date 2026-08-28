package main

import "fmt"

func switch_fun() {
	var day string

	if _, err := fmt.Scan(&day); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// switch day { // tagged switch
	// case "sat":
	// 	fmt.Println("Today Saturday")
	// case "sun":
	// 	fmt.Println("Today Sunday")
	// default:
	// 	fmt.Println("Today Off Day")
	// }

	switch { // normal switch
	case day == "sat":
		fmt.Println("Today Saturday")
	case day == "sun":
		fmt.Println("Today Sunday")
	default:
		fmt.Println("Today Off Day")
	}
}

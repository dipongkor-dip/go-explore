package main

import "fmt"

func conditions() {
	// fmt.Printf("Hello World!!\n")

	// age := 20

	var sc int

	fmt.Print("Enter your score: ")
	// _, err := fmt.Scan(&sc)

	if _, err := fmt.Scan(&sc); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Your score: %d\n", sc)

	if score := sc; score >= 80 {
		fmt.Println("You Got Gold Medal!!")
	} else if score >= 60 {
		fmt.Println("You Got Silver Medal!!")
	} else if score >= 50 {
		fmt.Println("You Got Bronze Medal!!")
	} else {
		fmt.Println("You Got participation certificate")
	}

	fmt.Printf("Your score: %d\n", sc)
}

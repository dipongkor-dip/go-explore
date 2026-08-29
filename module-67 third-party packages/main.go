package main

import (
	"fmt"
	"third-party-pack/callbackFunctions"

	"github.com/fatih/color"
)

// init functions parameters not allowed & return not allowed
func init() {
	fmt.Println("Welcome to init() function")
}

func init() {
	fmt.Println("Hello! init() function")
}

func main() {
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")
	color.Red("We have red")

	fmt.Println("Welcome to main() function")

	callbackFunctions.CallbackFunc()
}

// go mod tidy  == this command remove unnecessary imports
// The main() and init() functions are special functions in Go that control how a program starts and initializes.

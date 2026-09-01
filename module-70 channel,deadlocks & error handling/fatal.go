package main

import (
	"fmt"
	"log"
)

func fileHandling() {
	defer func() {
		fmt.Println("deferred function run")
	}()

	// not call any other function like defer func
	log.Fatal("Something very big has happened.")
}

func fetal_func() {
	fileHandling()
}

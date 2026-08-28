package main

import "fmt"

// ts --> interface for data type
// go --> interface behavior contract

type Physics struct{}

func (p Physics) read() {
	fmt.Println("Reading the Physics book......")
}

type Chemistry struct{}

func (c Chemistry) read() {
	fmt.Println("Reading the Chemistry book......")
}

type Biology struct {
	title string
}

func (b Biology) read() {
	fmt.Println("Reading the Biology book......", b.title)
}

type Book interface {
	read()
}

func checkBook(b Book) {
	b.read()
}

// oop pillars = abstraction, polymorphism, inheritance and encapsulation

func interface_func() {
	// phy := Physics{}
	// checkBook(phy)

	// che := Chemistry{}
	// checkBook(che)

	bio := Biology{title: "Haji"}
	checkBook(bio)
}

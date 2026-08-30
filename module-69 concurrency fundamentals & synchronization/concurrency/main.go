package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var start = time.Now()
	// sequential/blocking/one-by-one
	// uploadFile()
	// saveToDb() // waiting ...
	// sendEmail() // waiting ...

	// goroutine
	wg.Add(1)
	go uploadFile()

	wg.Add(1)
	go saveToDb()

	wg.Add(1)
	go sendEmail()

	wg.Wait() // waiting ... until counter is 0

	// time.Sleep(3 * time.Second)

	fmt.Println("✅ all tasks done")
	fmt.Println("🇹 time taken", time.Since(start))
}

func uploadFile() {
	fmt.Println("Uploading file ...")
	time.Sleep(3 * time.Second)

	fmt.Println("✅ File upload done!")
	// wg.Add(-1) // 2
	// alternative
	wg.Done()
}

func saveToDb() {
	fmt.Println("Saving file ...")
	time.Sleep(1 * time.Second)

	fmt.Println("✅ Saved done!")
	// wg.Add(-1) // 1
	wg.Done()
}

func sendEmail() {
	fmt.Println("Sending file ...")
	time.Sleep(2 * time.Second)

	fmt.Println("✅ Email sending done!")
	// wg.Add(-1) // 0
	wg.Done()
}

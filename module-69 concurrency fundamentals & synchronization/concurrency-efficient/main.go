package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var start = time.Now()

	// goroutine
	// wg.Add(1)
	// go uploadFiles()
	// alternative ~ inbuild Go using
	wg.Go(uploadFile)
	wg.Go(saveToDb)
	wg.Go(sendEmail)

	wg.Wait() // waiting ... until counter is 0

	fmt.Println("✅ all tasks done")
	fmt.Println("🇹 time taken", time.Since(start))
}

func uploadFile() {
	// defer wg.Done() // alternative inbuild Go using
	fmt.Println("Uploading file ...")
	time.Sleep(3 * time.Second)

	fmt.Println("✅ File upload done!")
	// wg.Done() // alternative line 29 ~ defer wg.Done()
}

func saveToDb() {
	// defer wg.Done()
	fmt.Println("Saving file ...")
	time.Sleep(1 * time.Second)

	fmt.Println("✅ Saved done!")
}

func sendEmail() {
	// defer wg.Done()
	fmt.Println("Sending file ...")
	time.Sleep(2 * time.Second)

	fmt.Println("✅ Email sending done!")
}

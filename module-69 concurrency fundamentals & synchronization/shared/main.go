package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var start = time.Now()

	wg.Go(uploadFile)
	wg.Go(saveToDb)
	wg.Go(sendEmail)

	wg.Wait() // waiting ... until counter is 0

	fmt.Println("✅ all tasks done")
	fmt.Println("🇹 time taken", time.Since(start))
}

func uploadFile() {
	fmt.Println("Uploading file ...")
	time.Sleep(3 * time.Second) // simulating file upload time

	fmt.Println("✅ File upload done!")

	// var fileUrl string = "https://s3/image.png"
	// return fileUrl
}

func saveToDb() {
	fmt.Println("Saving file ...")
	time.Sleep(1 * time.Second)

	fmt.Println("✅ Saved done!")
}

func sendEmail() {
	fmt.Println("Sending file ...")
	time.Sleep(2 * time.Second)

	fmt.Println("✅ Email sending done!")
}

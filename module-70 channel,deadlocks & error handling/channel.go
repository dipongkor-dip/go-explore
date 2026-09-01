package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func channel_func() {
	fmt.Println("Hello world!!")

	var channel = make(chan string) // un-buffered channel

	// go uploadFile(channel)

	// var fileUrl = <-channel // blocking ... (if data not return from uploadFile ~ deadlock err)
	// fmt.Println("data from channel", fileUrl)

	// catch deadlock by waitGroup ---------
	wg.Add(1)
	go uploadDoc(channel)

	var docFile = <-channel // blocking ... (if channel data not use ~ deadlock)
	wg.Wait()
	// var docFile = <-channel // deadlock

	fmt.Println("data from channel", docFile)

}

func uploadFile(c chan string) {
	fmt.Println("⌚ uploading file ...")
	time.Sleep(3 * time.Second)

	var fileUrl = "https://image.png"

	fmt.Println("✅ file upload done!")
	c <- fileUrl // data push channel
}

func uploadDoc(c chan string) {
	defer wg.Add(-1)

	fmt.Println("⌚ uploading file ...")
	time.Sleep(3 * time.Second)

	var docFile = "https://abc.pdf"

	fmt.Println("✅ file upload done!")
	c <- docFile // data push channel
}

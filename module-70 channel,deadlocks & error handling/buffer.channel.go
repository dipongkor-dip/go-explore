package main

import (
	"fmt"
	"time"
)

func buffer_channel_func() {
	var channel = make(chan string, 2) // capacity 2 of channel

	// 1st
	// upload file and get the url
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "file uploaded!"
	}()

	// 2nd
	// save file url
	go func() {
		time.Sleep(1 * time.Second)
		channel <- "file url saved!"
	}()

	// waiting if channel capacity is full, then push this operation
	// email sent
	go func() {
		time.Sleep(3 * time.Second)
		channel <- "email sent!"
	}()

	for range 3 {
		var data = <-channel
		fmt.Println(data)
	}
}

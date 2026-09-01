package main

import (
	"fmt"
	"time"
)

func behavior_channel_func() {
	var channel = make(chan string, 3)

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

	// 3rd
	// email sent
	go func() {
		time.Sleep(3 * time.Second)
		channel <- "email sent!"
	}()

	// when range is 4 this time waiting ... (deadlock err)
	for range 4 {
		var data = <-channel
		fmt.Println(data)
	}
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

/**
@ Concurrent website health checker
- concurrent execution using goroutines
- communication between goroutines using channels
- error handling with err != nil
- blocking behavior of channels
- A simple real-world concurrency pattern
*/

type Result struct {
	url    string
	status string
	err    error
}

func checkUrls(url string, ch chan Result) {
	var res, err = http.Get(url)

	if err != nil {
		fmt.Println(url, "is down")

		ch <- Result{url: url, status: "is down", err: err}

		return
	}

	defer res.Body.Close()

	ch <- Result{url: url, status: "is up and running", err: nil}
}

func main() {
	var urls = []string{"https://google.com", "https://github.com", "https://wrong-url-test.com"}

	var start = time.Now()

	var channel = make(chan Result)

	for _, url := range urls {
		go checkUrls(url, channel)
	}

	for range urls {
		var result = <-channel

		if result.err != nil {
			fmt.Println(result.url, result.status, "Error:", result.err)
			continue
		}

		fmt.Println(result)
	}

	fmt.Println("time taken", time.Since(start))

	fmt.Println("All urls checked successfully")
}

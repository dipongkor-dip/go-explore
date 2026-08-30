package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex // rac condition issue fixed by mutex

var counter int

func main() {
	for range 100 {
		wg.Go(increment)
	}

	wg.Wait()

	fmt.Println("counter value is", counter)
}

func increment() {
	mu.Lock()
	defer mu.Unlock() // when this function all tasks is completed then call
	counter += 1
	// mu.Unlock()
}

// go race condition

package main

import (
	"fmt"
	"sync"
)

var (
	// The shared resource that multiple goroutines will access
	counter int
	// The mutex to protect the counter
	mu sync.Mutex
	// WaitGroup to wait for all goroutines to finish
	wg sync.WaitGroup
)

func increment() {
	// Lock the mutex before accessing the critical section
	mu.Lock()
	// Ensure the mutex is unlocked when the function returns
	defer mu.Unlock()
	
	// Critical section: only one goroutine can execute this at a time
	counter++
	
	// Signal to the WaitGroup that this goroutine is done
	wg.Done()
}

func main() {
	numGoroutines := 1000

	// Add the number of goroutines to the WaitGroup counter
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go increment()
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// The final counter value will reliably be 1000
	fmt.Println("Final Counter:", counter)
}


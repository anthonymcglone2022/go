package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // Declare the WaitGroup
    var wg sync.WaitGroup

	// Create several goroutines
    for i := 1; i <= 6; i++ {
	
		// Increment the counter
        wg.Add(1)

        i := i

        go func() {
			// Decrement the counter when the goroutine is finished
            defer wg.Done()
            goRoutine(i)
        }()
    }

	// Wait for the counter to go back to 0
    wg.Wait()
}

func goRoutine(id int) {
	fmt.Println(id, "Routine, started\n")
    time.Sleep(time.Second)
	fmt.Println(id, "Routine, FINISHED\n")
}
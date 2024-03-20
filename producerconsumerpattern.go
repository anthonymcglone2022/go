package main

import (
	"fmt"
)

// consumer function to square numbers 
// The argument queue <-chan int indicates to only ever receive on the queue channel
// The argument results chan<- string  indicates to only ever send to the results channel

func consumer(queue <-chan int, results chan<- string, workerId int) {

    // The consumer receives an item from the queue
    for n := range queue {
        fmt.Println("Worker ID:", workerId, "taking queue job number", n)
		
        // Send the square to the results channel
        results <- fmt.Sprintf("The square of %d is %d", n, square(n))
		
        fmt.Println("Worker ID:", workerId, "finished queue job number", n)
    }
}

func square(n int) int {
    return n * n
}

func main() {
    // Buffered channel for the jobs queue
    queue := make(chan int, 9)
	
    // Store the square numbers in a separate channel
    results := make(chan string, 9)

    // Consumer pool
    go consumer(queue, results, 1)
    go consumer(queue, results, 2)
    go consumer(queue, results, 3)

    // Producer
    for i := 0; i < 9; i++ {
        queue <- i
    }
    close(queue)

    // Print the squares
    for j := 0; j < 9; j++ {

        // Receive from the results channel
        fmt.Println(<-results)
    }
}
